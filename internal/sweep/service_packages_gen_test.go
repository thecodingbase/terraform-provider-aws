// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sweep_test

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/service/accessanalyzer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/account"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/acmpca"
	"github.com/hashicorp/terraform-provider-aws/internal/service/amp"
	"github.com/hashicorp/terraform-provider-aws/internal/service/amplify"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apigatewayv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appautoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appfabric"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appintegrations"
	"github.com/hashicorp/terraform-provider-aws/internal/service/applicationinsights"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	"github.com/hashicorp/terraform-provider-aws/internal/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/internal/service/appsync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/athena"
	"github.com/hashicorp/terraform-provider-aws/internal/service/auditmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/service/autoscalingplans"
	"github.com/hashicorp/terraform-provider-aws/internal/service/backup"
	"github.com/hashicorp/terraform-provider-aws/internal/service/batch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/bcmdataexports"
	"github.com/hashicorp/terraform-provider-aws/internal/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/internal/service/bedrockagent"
	"github.com/hashicorp/terraform-provider-aws/internal/service/budgets"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ce"
	"github.com/hashicorp/terraform-provider-aws/internal/service/chatbot"
	"github.com/hashicorp/terraform-provider-aws/internal/service/chime"
	"github.com/hashicorp/terraform-provider-aws/internal/service/chimesdkmediapipelines"
	"github.com/hashicorp/terraform-provider-aws/internal/service/chimesdkvoice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cleanrooms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloud9"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudfront"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudfrontkeyvaluestore"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudhsmv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudtrail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cloudwatch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codeartifact"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codecatalyst"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codecommit"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codeguruprofiler"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codegurureviewer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codepipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarconnections"
	"github.com/hashicorp/terraform-provider-aws/internal/service/codestarnotifications"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidentity"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidp"
	"github.com/hashicorp/terraform-provider-aws/internal/service/comprehend"
	"github.com/hashicorp/terraform-provider-aws/internal/service/computeoptimizer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/configservice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/connect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/connectcases"
	"github.com/hashicorp/terraform-provider-aws/internal/service/controltower"
	"github.com/hashicorp/terraform-provider-aws/internal/service/costoptimizationhub"
	"github.com/hashicorp/terraform-provider-aws/internal/service/cur"
	"github.com/hashicorp/terraform-provider-aws/internal/service/customerprofiles"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dataexchange"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datapipeline"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datasync"
	"github.com/hashicorp/terraform-provider-aws/internal/service/datazone"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dax"
	"github.com/hashicorp/terraform-provider-aws/internal/service/deploy"
	"github.com/hashicorp/terraform-provider-aws/internal/service/detective"
	"github.com/hashicorp/terraform-provider-aws/internal/service/devicefarm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/devopsguru"
	"github.com/hashicorp/terraform-provider-aws/internal/service/directconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dlm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/docdb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/docdbelastic"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/dynamodb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecrpublic"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ecs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/efs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/eks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elasticsearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elastictranscoder"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/elbv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emr"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrcontainers"
	"github.com/hashicorp/terraform-provider-aws/internal/service/emrserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/events"
	"github.com/hashicorp/terraform-provider-aws/internal/service/evidently"
	"github.com/hashicorp/terraform-provider-aws/internal/service/finspace"
	"github.com/hashicorp/terraform-provider-aws/internal/service/firehose"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/fsx"
	"github.com/hashicorp/terraform-provider-aws/internal/service/gamelift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glacier"
	"github.com/hashicorp/terraform-provider-aws/internal/service/globalaccelerator"
	"github.com/hashicorp/terraform-provider-aws/internal/service/glue"
	"github.com/hashicorp/terraform-provider-aws/internal/service/grafana"
	"github.com/hashicorp/terraform-provider-aws/internal/service/greengrass"
	"github.com/hashicorp/terraform-provider-aws/internal/service/groundstation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/service/healthlake"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iam"
	"github.com/hashicorp/terraform-provider-aws/internal/service/identitystore"
	"github.com/hashicorp/terraform-provider-aws/internal/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/internal/service/inspector"
	"github.com/hashicorp/terraform-provider-aws/internal/service/inspector2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/internetmonitor"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iot"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iotanalytics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/iotevents"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ivs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ivschat"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafka"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kafkaconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kendra"
	"github.com/hashicorp/terraform-provider-aws/internal/service/keyspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalytics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalyticsv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kinesisvideo"
	"github.com/hashicorp/terraform-provider-aws/internal/service/kms"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lakeformation"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lambda"
	"github.com/hashicorp/terraform-provider-aws/internal/service/launchwizard"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lexmodels"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lexv2models"
	"github.com/hashicorp/terraform-provider-aws/internal/service/licensemanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lightsail"
	"github.com/hashicorp/terraform-provider-aws/internal/service/location"
	"github.com/hashicorp/terraform-provider-aws/internal/service/logs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/lookoutmetrics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/m2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/macie2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediaconnect"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediaconvert"
	"github.com/hashicorp/terraform-provider-aws/internal/service/medialive"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediapackage"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediapackagev2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mediastore"
	"github.com/hashicorp/terraform-provider-aws/internal/service/memorydb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/meta"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mq"
	"github.com/hashicorp/terraform-provider-aws/internal/service/mwaa"
	"github.com/hashicorp/terraform-provider-aws/internal/service/neptune"
	"github.com/hashicorp/terraform-provider-aws/internal/service/neptunegraph"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkfirewall"
	"github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/oam"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearch"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/opsworks"
	"github.com/hashicorp/terraform-provider-aws/internal/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/internal/service/osis"
	"github.com/hashicorp/terraform-provider-aws/internal/service/outposts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/paymentcryptography"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pcaconnectorad"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pinpoint"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pipes"
	"github.com/hashicorp/terraform-provider-aws/internal/service/polly"
	"github.com/hashicorp/terraform-provider-aws/internal/service/pricing"
	"github.com/hashicorp/terraform-provider-aws/internal/service/qbusiness"
	"github.com/hashicorp/terraform-provider-aws/internal/service/qldb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ram"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rbin"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftdata"
	"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rekognition"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourceexplorer2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourcegroups"
	"github.com/hashicorp/terraform-provider-aws/internal/service/resourcegroupstaggingapi"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rolesanywhere"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53domains"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53profiles"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53recoverycontrolconfig"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53recoveryreadiness"
	"github.com/hashicorp/terraform-provider-aws/internal/service/route53resolver"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rum"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3control"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3outposts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/service/scheduler"
	"github.com/hashicorp/terraform-provider-aws/internal/service/schemas"
	"github.com/hashicorp/terraform-provider-aws/internal/service/secretsmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/service/securityhub"
	"github.com/hashicorp/terraform-provider-aws/internal/service/securitylake"
	"github.com/hashicorp/terraform-provider-aws/internal/service/serverlessrepo"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalogappregistry"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicediscovery"
	"github.com/hashicorp/terraform-provider-aws/internal/service/servicequotas"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ses"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sfn"
	"github.com/hashicorp/terraform-provider-aws/internal/service/shield"
	"github.com/hashicorp/terraform-provider-aws/internal/service/signer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/simpledb"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sns"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sqs"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssmcontacts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssmincidents"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssmsap"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sso"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/internal/service/storagegateway"
	"github.com/hashicorp/terraform-provider-aws/internal/service/sts"
	"github.com/hashicorp/terraform-provider-aws/internal/service/swf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/synthetics"
	"github.com/hashicorp/terraform-provider-aws/internal/service/timestreamwrite"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transcribe"
	"github.com/hashicorp/terraform-provider-aws/internal/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/internal/service/verifiedpermissions"
	"github.com/hashicorp/terraform-provider-aws/internal/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/service/waf"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafregional"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wafv2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/wellarchitected"
	"github.com/hashicorp/terraform-provider-aws/internal/service/worklink"
	"github.com/hashicorp/terraform-provider-aws/internal/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/service/xray"
)

func servicePackages(ctx context.Context) []conns.ServicePackage {
	v := []conns.ServicePackage{
		accessanalyzer.ServicePackage(ctx),
		account.ServicePackage(ctx),
		acm.ServicePackage(ctx),
		acmpca.ServicePackage(ctx),
		amp.ServicePackage(ctx),
		amplify.ServicePackage(ctx),
		apigateway.ServicePackage(ctx),
		apigatewayv2.ServicePackage(ctx),
		appautoscaling.ServicePackage(ctx),
		appconfig.ServicePackage(ctx),
		appfabric.ServicePackage(ctx),
		appflow.ServicePackage(ctx),
		appintegrations.ServicePackage(ctx),
		applicationinsights.ServicePackage(ctx),
		appmesh.ServicePackage(ctx),
		apprunner.ServicePackage(ctx),
		appstream.ServicePackage(ctx),
		appsync.ServicePackage(ctx),
		athena.ServicePackage(ctx),
		auditmanager.ServicePackage(ctx),
		autoscaling.ServicePackage(ctx),
		autoscalingplans.ServicePackage(ctx),
		backup.ServicePackage(ctx),
		batch.ServicePackage(ctx),
		bcmdataexports.ServicePackage(ctx),
		bedrock.ServicePackage(ctx),
		bedrockagent.ServicePackage(ctx),
		budgets.ServicePackage(ctx),
		ce.ServicePackage(ctx),
		chatbot.ServicePackage(ctx),
		chime.ServicePackage(ctx),
		chimesdkmediapipelines.ServicePackage(ctx),
		chimesdkvoice.ServicePackage(ctx),
		cleanrooms.ServicePackage(ctx),
		cloud9.ServicePackage(ctx),
		cloudcontrol.ServicePackage(ctx),
		cloudformation.ServicePackage(ctx),
		cloudfront.ServicePackage(ctx),
		cloudfrontkeyvaluestore.ServicePackage(ctx),
		cloudhsmv2.ServicePackage(ctx),
		cloudsearch.ServicePackage(ctx),
		cloudtrail.ServicePackage(ctx),
		cloudwatch.ServicePackage(ctx),
		codeartifact.ServicePackage(ctx),
		codebuild.ServicePackage(ctx),
		codecatalyst.ServicePackage(ctx),
		codecommit.ServicePackage(ctx),
		codeguruprofiler.ServicePackage(ctx),
		codegurureviewer.ServicePackage(ctx),
		codepipeline.ServicePackage(ctx),
		codestarconnections.ServicePackage(ctx),
		codestarnotifications.ServicePackage(ctx),
		cognitoidentity.ServicePackage(ctx),
		cognitoidp.ServicePackage(ctx),
		comprehend.ServicePackage(ctx),
		computeoptimizer.ServicePackage(ctx),
		configservice.ServicePackage(ctx),
		connect.ServicePackage(ctx),
		connectcases.ServicePackage(ctx),
		controltower.ServicePackage(ctx),
		costoptimizationhub.ServicePackage(ctx),
		cur.ServicePackage(ctx),
		customerprofiles.ServicePackage(ctx),
		dataexchange.ServicePackage(ctx),
		datapipeline.ServicePackage(ctx),
		datasync.ServicePackage(ctx),
		datazone.ServicePackage(ctx),
		dax.ServicePackage(ctx),
		deploy.ServicePackage(ctx),
		detective.ServicePackage(ctx),
		devicefarm.ServicePackage(ctx),
		devopsguru.ServicePackage(ctx),
		directconnect.ServicePackage(ctx),
		dlm.ServicePackage(ctx),
		dms.ServicePackage(ctx),
		docdb.ServicePackage(ctx),
		docdbelastic.ServicePackage(ctx),
		ds.ServicePackage(ctx),
		dynamodb.ServicePackage(ctx),
		ec2.ServicePackage(ctx),
		ecr.ServicePackage(ctx),
		ecrpublic.ServicePackage(ctx),
		ecs.ServicePackage(ctx),
		efs.ServicePackage(ctx),
		eks.ServicePackage(ctx),
		elasticache.ServicePackage(ctx),
		elasticbeanstalk.ServicePackage(ctx),
		elasticsearch.ServicePackage(ctx),
		elastictranscoder.ServicePackage(ctx),
		elb.ServicePackage(ctx),
		elbv2.ServicePackage(ctx),
		emr.ServicePackage(ctx),
		emrcontainers.ServicePackage(ctx),
		emrserverless.ServicePackage(ctx),
		events.ServicePackage(ctx),
		evidently.ServicePackage(ctx),
		finspace.ServicePackage(ctx),
		firehose.ServicePackage(ctx),
		fis.ServicePackage(ctx),
		fms.ServicePackage(ctx),
		fsx.ServicePackage(ctx),
		gamelift.ServicePackage(ctx),
		glacier.ServicePackage(ctx),
		globalaccelerator.ServicePackage(ctx),
		glue.ServicePackage(ctx),
		grafana.ServicePackage(ctx),
		greengrass.ServicePackage(ctx),
		groundstation.ServicePackage(ctx),
		guardduty.ServicePackage(ctx),
		healthlake.ServicePackage(ctx),
		iam.ServicePackage(ctx),
		identitystore.ServicePackage(ctx),
		imagebuilder.ServicePackage(ctx),
		inspector.ServicePackage(ctx),
		inspector2.ServicePackage(ctx),
		internetmonitor.ServicePackage(ctx),
		iot.ServicePackage(ctx),
		iotanalytics.ServicePackage(ctx),
		iotevents.ServicePackage(ctx),
		ivs.ServicePackage(ctx),
		ivschat.ServicePackage(ctx),
		kafka.ServicePackage(ctx),
		kafkaconnect.ServicePackage(ctx),
		kendra.ServicePackage(ctx),
		keyspaces.ServicePackage(ctx),
		kinesis.ServicePackage(ctx),
		kinesisanalytics.ServicePackage(ctx),
		kinesisanalyticsv2.ServicePackage(ctx),
		kinesisvideo.ServicePackage(ctx),
		kms.ServicePackage(ctx),
		lakeformation.ServicePackage(ctx),
		lambda.ServicePackage(ctx),
		launchwizard.ServicePackage(ctx),
		lexmodels.ServicePackage(ctx),
		lexv2models.ServicePackage(ctx),
		licensemanager.ServicePackage(ctx),
		lightsail.ServicePackage(ctx),
		location.ServicePackage(ctx),
		logs.ServicePackage(ctx),
		lookoutmetrics.ServicePackage(ctx),
		m2.ServicePackage(ctx),
		macie2.ServicePackage(ctx),
		mediaconnect.ServicePackage(ctx),
		mediaconvert.ServicePackage(ctx),
		medialive.ServicePackage(ctx),
		mediapackage.ServicePackage(ctx),
		mediapackagev2.ServicePackage(ctx),
		mediastore.ServicePackage(ctx),
		memorydb.ServicePackage(ctx),
		meta.ServicePackage(ctx),
		mq.ServicePackage(ctx),
		mwaa.ServicePackage(ctx),
		neptune.ServicePackage(ctx),
		neptunegraph.ServicePackage(ctx),
		networkfirewall.ServicePackage(ctx),
		networkmanager.ServicePackage(ctx),
		oam.ServicePackage(ctx),
		opensearch.ServicePackage(ctx),
		opensearchserverless.ServicePackage(ctx),
		opsworks.ServicePackage(ctx),
		organizations.ServicePackage(ctx),
		osis.ServicePackage(ctx),
		outposts.ServicePackage(ctx),
		paymentcryptography.ServicePackage(ctx),
		pcaconnectorad.ServicePackage(ctx),
		pinpoint.ServicePackage(ctx),
		pipes.ServicePackage(ctx),
		polly.ServicePackage(ctx),
		pricing.ServicePackage(ctx),
		qbusiness.ServicePackage(ctx),
		qldb.ServicePackage(ctx),
		quicksight.ServicePackage(ctx),
		ram.ServicePackage(ctx),
		rbin.ServicePackage(ctx),
		rds.ServicePackage(ctx),
		redshift.ServicePackage(ctx),
		redshiftdata.ServicePackage(ctx),
		redshiftserverless.ServicePackage(ctx),
		rekognition.ServicePackage(ctx),
		resourceexplorer2.ServicePackage(ctx),
		resourcegroups.ServicePackage(ctx),
		resourcegroupstaggingapi.ServicePackage(ctx),
		rolesanywhere.ServicePackage(ctx),
		route53.ServicePackage(ctx),
		route53domains.ServicePackage(ctx),
		route53profiles.ServicePackage(ctx),
		route53recoverycontrolconfig.ServicePackage(ctx),
		route53recoveryreadiness.ServicePackage(ctx),
		route53resolver.ServicePackage(ctx),
		rum.ServicePackage(ctx),
		s3.ServicePackage(ctx),
		s3control.ServicePackage(ctx),
		s3outposts.ServicePackage(ctx),
		sagemaker.ServicePackage(ctx),
		scheduler.ServicePackage(ctx),
		schemas.ServicePackage(ctx),
		secretsmanager.ServicePackage(ctx),
		securityhub.ServicePackage(ctx),
		securitylake.ServicePackage(ctx),
		serverlessrepo.ServicePackage(ctx),
		servicecatalog.ServicePackage(ctx),
		servicecatalogappregistry.ServicePackage(ctx),
		servicediscovery.ServicePackage(ctx),
		servicequotas.ServicePackage(ctx),
		ses.ServicePackage(ctx),
		sesv2.ServicePackage(ctx),
		sfn.ServicePackage(ctx),
		shield.ServicePackage(ctx),
		signer.ServicePackage(ctx),
		simpledb.ServicePackage(ctx),
		sns.ServicePackage(ctx),
		sqs.ServicePackage(ctx),
		ssm.ServicePackage(ctx),
		ssmcontacts.ServicePackage(ctx),
		ssmincidents.ServicePackage(ctx),
		ssmsap.ServicePackage(ctx),
		sso.ServicePackage(ctx),
		ssoadmin.ServicePackage(ctx),
		storagegateway.ServicePackage(ctx),
		sts.ServicePackage(ctx),
		swf.ServicePackage(ctx),
		synthetics.ServicePackage(ctx),
		timestreamwrite.ServicePackage(ctx),
		transcribe.ServicePackage(ctx),
		transfer.ServicePackage(ctx),
		verifiedpermissions.ServicePackage(ctx),
		vpclattice.ServicePackage(ctx),
		waf.ServicePackage(ctx),
		wafregional.ServicePackage(ctx),
		wafv2.ServicePackage(ctx),
		wellarchitected.ServicePackage(ctx),
		worklink.ServicePackage(ctx),
		workspaces.ServicePackage(ctx),
		xray.ServicePackage(ctx),
	}

	return slices.Clone(v)
}
