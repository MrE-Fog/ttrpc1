package mgrpc

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type mgrpcGenerator struct {
	*generator.Generator
	generator.PluginImports

	typeurlPkg generator.Single
	mgrpcPkg   generator.Single
	contextPkg generator.Single
}

func init() {
	generator.RegisterPlugin(new(mgrpcGenerator))
}

func (p *mgrpcGenerator) Name() string {
	return "mgrpc"
}

func (p *mgrpcGenerator) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *mgrpcGenerator) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.contextPkg = p.NewImport("context")
	p.typeurlPkg = p.NewImport("github.com/containerd/typeurl")
	p.mgrpcPkg = p.NewImport("github.com/stevvooe/mgrpc")

	for _, service := range file.GetService() {
		serviceName := service.GetName()
		if pkg := file.GetPackage(); pkg != "" {
			serviceName = pkg + "." + serviceName
		}

		p.genService(serviceName, service)
	}
}

func (p *mgrpcGenerator) genService(fullName string, service *descriptor.ServiceDescriptorProto) {
	serviceName := service.GetName() + "Service"
	p.P()
	p.P("type ", serviceName, " interface{")
	p.In()
	for _, method := range service.Method {
		p.P(method.GetName(),
			"(ctx ", p.contextPkg.Use(), ".Context, ",
			"req *", p.typeName(method.GetInputType()), ") ",
			"(*", p.typeName(method.GetOutputType()), ", error)")

	}
	p.Out()
	p.P("}")

	p.P()
	// registration method
	p.P("func Register", serviceName, "(srv *", p.mgrpcPkg.Use(), ".Server, svc ", serviceName, ") error {")
	p.In()
	p.P(`return srv.Register("`, fullName, `", map[string]`, p.mgrpcPkg.Use(), ".Handler{")
	p.In()
	for _, method := range service.Method {
		p.P(`"`, method.GetName(), `": `, p.mgrpcPkg.Use(), `.HandlerFunc(func(ctx context.Context, req interface{}) (interface{}, error) {`)
		p.In()
		p.P("return svc.", method.GetName(), "(ctx, req.(*", p.typeName(method.GetInputType()), "))")
		p.Out()
		p.P("}),")
	}
	p.Out()
	p.P("})")
	p.Out()
	p.P("}")

	clientType := service.GetName() + "Client"
	p.P()
	p.P("type ", clientType, " struct{")
	p.In()
	p.P("client *", p.mgrpcPkg.Use(), ".Client")
	p.Out()
	p.P("}")
	for _, method := range service.Method {
		p.P()
		p.P("func (c *", clientType, ") ", method.GetName(),
			"(ctx ", p.contextPkg.Use(), ".Context, ",
			"req *", p.typeName(method.GetInputType()), ") ",
			"(*", p.typeName(method.GetOutputType()), ", error) {")
		p.In()
		p.P("resp, err := c.client.Call(ctx, ", `"`+fullName+`", `, `"`+method.GetName()+`"`, ", req)")
		p.P("if err != nil {")
		p.In()
		p.P("return nil, err")
		p.Out()
		p.P("}")
		p.P("return resp.(*", p.typeName(method.GetOutputType()), "), nil")
		p.Out()
		p.P("}")
	}
}

func (p *mgrpcGenerator) objectNamed(name string) generator.Object {
	p.Generator.RecordTypeUse(name)
	return p.Generator.ObjectNamed(name)
}

func (p *mgrpcGenerator) typeName(str string) string {
	return p.Generator.TypeName(p.objectNamed(str))
}
