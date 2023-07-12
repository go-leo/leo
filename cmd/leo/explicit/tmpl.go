package explicit

import _ "embed"

var (
	//go:embed tmpl/cmd_wire.go.template
	cmdWireContent string

	//go:embed tmpl/cmd_main.go.template
	cmdMainContent string

	//go:embed tmpl/cmd_text_config.yaml.template
	cmdTextConfigContent string

	//go:embed tmpl/cmd_nacos_config.yaml.template
	cmdNacosConfigContent string
)

//go:embed tmpl/app_root_wire.go.template
var appRootWireContent string

var (
	//go:embed tmpl/presentation_wire.go.template
	presentationWireContent string

	//go:embed tmpl/presentation_assemblers.go.template
	presentationAssemblersContent string

	//go:embed tmpl/presentation_assembler_wire.go.template
	presentationAssemblerWireContent string

	//go:embed tmpl/presentation_runner_hello.go.template
	presentationRunnerHelloContent string

	//go:embed tmpl/presentation_runner_runners.go.template
	presentationRunnerRunnersContent string

	//go:embed tmpl/presentation_runner_wire.go.template
	presentationRunnerWireContent string

	//go:embed tmpl/presentation_controller_hello.go.template
	presentationControllerHelloContent string

	//go:embed tmpl/presentation_controller_wire.go.template
	presentationControllerWireContent string

	//go:embed tmpl/presentation_provider_hello.go.template
	presentationProviderHelloContent string

	//go:embed tmpl/presentation_provider_wire.go.template
	presentationProviderWireContent string

	//go:embed tmpl/presentation_task_hello.go.template
	presentationTaskHelloContent string

	//go:embed tmpl/presentation_task_tasks.go.template
	presentationTaskTasksContent string

	//go:embed tmpl/presentation_task_wire.go.template
	presentationTaskWireContent string
)

var (
	//go:embed tmpl/application_command_commands.go.template
	applicationCommandCommandsContent string

	//go:embed tmpl/application_command_wire.go.template
	applicationCommandWireContent string

	//go:embed tmpl/application_query_queries.go.template
	applicationQueryQueriesContent string

	//go:embed tmpl/application_query_wire.go.template
	applicationQueryWireContent string

	//go:embed tmpl/application_wire.go.template
	applicationWireContent string
)

//go:embed tmpl/domain_wire.go.template
var domainWireContent string

var (
	//go:embed tmpl/infrastructure_wire.go.template
	infrastructureContent string

	//go:embed tmpl/infrastructure_client_wire.go.template
	infrastructureClientWireContent string

	//go:embed tmpl/infrastructure_publisher_wire.go.template
	infrastructurePublisherWireContent string

	//go:embed tmpl/infrastructure_repository_wire.go.template
	infrastructureRepositoryWireContent string

	//go:embed tmpl/infrastructure_converters.go.template
	infrastructureConvertersContent string

	//go:embed tmpl/infrastructure_converter_wire.go.template
	infrastructureConverterWireContent string
)

var (
	//go:embed tmpl/pkg_wire.go.template
	pkgWireContent string

	//go:embed tmpl/pkg_actuatorx_config.go.template
	pkgActuatorxConfigContent string

	//go:embed tmpl/pkg_configx_configuration.go.template
	pkgConfigxConfigurationContent string

	//go:embed tmpl/pkg_configx_load.go.template
	pkgConfigxLoadContent string

	//go:embed tmpl/pkg_configx_wire.go.template
	pkgConfigxWireContent string

	//go:embed tmpl/pkg_ginx_config.go.template
	pkgGinxConfigContent string

	//go:embed tmpl/pkg_ginx_middleware.go.template
	pkgGinxMiddlewareContent string

	//go:embed tmpl/pkg_grpcx_client.go.template
	pkggRPCxClientContent string

	//go:embed tmpl/pkg_grpcx_server.go.template
	pkggRPCxServerContent string

	//go:embed tmpl/pkg_grpcx_wire.go.template
	pkggRPCxWireContent string

	//go:embed tmpl/pkg_streamx_amqpx_amqp.go.template
	pkgStreamxAMQPxAMQPContent string

	//go:embed tmpl/pkg_streamx_kafkax_kafka.go.template
	pkgStreamxKafkaxKafkaContent string

	//go:embed tmpl/pkg_streamx_wire.go.template
	pkgStreamxWireContent string
)

var (
	//go:embed tmpl/scripts_shell_format.sh.template
	scriptsShellFormatContent string

	//go:embed tmpl/scripts_shell_gen.sh.template
	scriptsShellGenContent string

	//go:embed tmpl/scripts_shell_lint.sh.template
	scriptsShellLintContent string

	//go:embed tmpl/scripts_shell_protoc.sh.template
	scriptsShellProtocContent string

	//go:embed tmpl/scripts_shell_tools.sh.template
	scriptsShellToolsContent string

	//go:embed tmpl/scripts_shell_wire.sh.template
	scriptsShellWireContent string
)

//go:embed tmpl/doc.go.template
var docContent string

//go:embed tmpl/sample_wire.go.template
var sampleWireContent string

//go:embed tmpl/tools_wire.go.template
var toolsWireContent string

//go:embed tmpl/Makefile.template
var _MakefileContent string

var (
	//go:embed tmpl/api_http_hello.go.template
	apiHttpHelloContent string

	//go:embed tmpl/api_grpc_hello.proto.template
	apiGrpcHelloContent string
)
