# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .allow_list import *
from .api_oidc_config import *
from .ca_cert import *
from .cluster import *
from .cmek import *
from .database import *
from .finalize_version_upgrade import *
from .folder import *
from .get_cluster_cert import *
from .get_cockroach_cluster import *
from .get_connection_string import *
from .get_organization import *
from .get_person_user import *
from .log_export_config import *
from .maintenance_window import *
from .metric_export_cloudwatch_config import *
from .metric_export_datadog_config import *
from .private_endpoint_connection import *
from .private_endpoint_services import *
from .private_endpoint_trusted_owner import *
from .provider import *
from .sql_user import *
from .user_role_grants import *
from .version_deferral import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_cockroach.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumiverse_cockroach.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "cockroach",
  "mod": "index/allowList",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/allowList:AllowList": "AllowList"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/apiOidcConfig",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/apiOidcConfig:ApiOidcConfig": "ApiOidcConfig"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/caCert",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/caCert:CaCert": "CaCert"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/cluster",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/cluster:Cluster": "Cluster"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/cmek",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/cmek:Cmek": "Cmek"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/database",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/database:Database": "Database"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/finalizeVersionUpgrade",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/finalizeVersionUpgrade:FinalizeVersionUpgrade": "FinalizeVersionUpgrade"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/folder",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/folder:Folder": "Folder"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/logExportConfig",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/logExportConfig:LogExportConfig": "LogExportConfig"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/maintenanceWindow",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/maintenanceWindow:MaintenanceWindow": "MaintenanceWindow"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/metricExportCloudwatchConfig",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig": "MetricExportCloudwatchConfig"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/metricExportDatadogConfig",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/metricExportDatadogConfig:MetricExportDatadogConfig": "MetricExportDatadogConfig"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/privateEndpointConnection",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/privateEndpointConnection:PrivateEndpointConnection": "PrivateEndpointConnection"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/privateEndpointServices",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/privateEndpointServices:PrivateEndpointServices": "PrivateEndpointServices"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/privateEndpointTrustedOwner",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner": "PrivateEndpointTrustedOwner"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/sqlUser",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/sqlUser:SqlUser": "SqlUser"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/userRoleGrants",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/userRoleGrants:UserRoleGrants": "UserRoleGrants"
  }
 },
 {
  "pkg": "cockroach",
  "mod": "index/versionDeferral",
  "fqn": "pulumiverse_cockroach",
  "classes": {
   "cockroach:index/versionDeferral:VersionDeferral": "VersionDeferral"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "cockroach",
  "token": "pulumi:providers:cockroach",
  "fqn": "pulumiverse_cockroach",
  "class": "Provider"
 }
]
"""
)