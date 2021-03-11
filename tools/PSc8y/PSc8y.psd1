#
# Module manifest for module 'PSc8y'
#
# Generated by: Reuben Miller
#
# Generated on: 2019.10.16
#

@{

# Script module or binary module file associated with this manifest.
RootModule = 'PSc8y.psm1'

# Version number of this module.
ModuleVersion = '1.11.0'

# Supported PSEditions
CompatiblePSEditions = @(
    "Desktop",
    "Core"
)

# ID used to uniquely identify this module
GUID = '9b4ebc35-bd61-4d07-a528-81c5733d0732'

# Author of this module
Author = 'Reuben Miller'


# Company or vendor of this module
CompanyName = ''

# Copyright statement for this module
Copyright = '(c) 2021. All rights reserved.'

# Description of the functionality provided by this module
Description = 'Cumulocity REST API'

# Minimum version of the Windows PowerShell engine required by this module
PowerShellVersion = '5.1'

# Name of the Windows PowerShell host required by this module
# PowerShellHostName = ''

# Minimum version of the Windows PowerShell host required by this module
# PowerShellHostVersion = ''

# Minimum version of Microsoft .NET Framework required by this module. This prerequisite is valid for the PowerShell Desktop edition only.
# DotNetFrameworkVersion = '2.0'

# Minimum version of the common language runtime (CLR) required by this module. This prerequisite is valid for the PowerShell Desktop edition only.
# CLRVersion = ''

# Processor architecture (None, X86, Amd64) required by this module
# ProcessorArchitecture = ''

# Modules that must be imported into the global environment prior to importing this module
RequiredModules = @()

# Format files (.ps1xml) to be loaded when importing this module
FormatsToProcess = @(
	'format-data/agents.ps1xml',
    'format-data/alarms.ps1xml',
    'format-data/applicationReferences.ps1xml',
    'format-data/applications.ps1xml',
	'format-data/auditRecords.ps1xml',
	'format-data/bulkOperations.ps1xml',
    'format-data/deviceCredentials.ps1xml',
    'format-data/deviceGroups.ps1xml',
    'format-data/devices.ps1xml',
    'format-data/events.ps1xml',
    'format-data/identity.ps1xml',
    'format-data/inventoryRoles.ps1xml',
    'format-data/managedObjectReferences.ps1xml',
    'format-data/managedObjects.ps1xml',
    'format-data/measurements.ps1xml',
    'format-data/microservices.ps1xml',
    'format-data/operations.ps1xml',
    'format-data/options.ps1xml',
    'format-data/retentionRules.ps1xml',
    'format-data/roleReferences.ps1xml',
    'format-data/roles.ps1xml',
    'format-data/session.ps1xml',
    'format-data/tenants.ps1xml',
    'format-data/tenantStatistics.ps1xml',
    'format-data/userGroupMembership.ps1xml',
    'format-data/userGroups.ps1xml',
    'format-data/userReferences.ps1xml',
    'format-data/users.ps1xml'
)

# Functions to export from this module, for best performance, do not use wildcards and do not delete the entry, use an empty array if there are no functions to export.
FunctionsToExport = @(
	'Add-AssetToGroup',
	'Add-ChildAddition',
	'Add-ChildDeviceToDevice',
	'Add-ChildGroupToGroup',
	'Add-DeviceToGroup',
	'Add-RoleToGroup',
	'Add-RoleToUser',
	'Add-UserToGroup',
	'Approve-DeviceRequest',
	'Copy-Application',
	'Disable-Application',
	'Disable-Microservice',
	'Enable-Application',
	'Enable-Microservice',
	'Get-Agent',
	'Get-Alarm',
	'Get-AlarmCollection',
	'Get-AllTenantUsageSummaryStatistics',
	'Get-Application',
	'Get-ApplicationBinaryCollection',
	'Get-ApplicationCollection',
	'Get-ApplicationReferenceCollection',
	'Get-AuditRecord',
	'Get-AuditRecordCollection',
	'Get-Binary',
	'Get-BinaryCollection',
	'Get-BulkOperation',
	'Get-BulkOperationCollection',
	'Get-ChildAdditionCollection',
	'Get-ChildAssetCollection',
	'Get-ChildAssetReference',
	'Get-ChildDeviceCollection',
	'Get-ChildDeviceReference',
	'Get-CurrentApplication',
	'Get-CurrentApplicationSubscription',
	'Get-CurrentTenant',
	'Get-CurrentUser',
	'Get-CurrentUserInventoryRole',
	'Get-CurrentUserInventoryRoleCollection',
	'Get-DataBrokerConnector',
	'Get-DataBrokerConnectorCollection',
	'Get-Device',
	'Get-DeviceGroup',
	'Get-DeviceRequest',
	'Get-DeviceRequestCollection',
	'Get-Event',
	'Get-EventBinary',
	'Get-EventCollection',
	'Get-ExternalId',
	'Get-ExternalIdCollection',
	'Get-ManagedObject',
	'Get-ManagedObjectCollection',
	'Get-Measurement',
	'Get-MeasurementCollection',
	'Get-MeasurementSeries',
	'Get-Microservice',
	'Get-MicroserviceBootstrapUser',
	'Get-MicroserviceCollection',
	'Get-Operation',
	'Get-OperationCollection',
	'Get-RetentionRule',
	'Get-RetentionRuleCollection',
	'Get-RoleCollection',
	'Get-RoleReferenceCollectionFromGroup',
	'Get-RoleReferenceCollectionFromUser',
	'Get-SupportedMeasurements',
	'Get-SupportedOperations',
	'Get-SupportedSeries',
	'Get-SystemOption',
	'Get-SystemOptionCollection',
	'Get-Tenant',
	'Get-TenantCollection',
	'Get-TenantOption',
	'Get-TenantOptionCollection',
	'Get-TenantOptionForCategory',
	'Get-TenantStatisticsCollection',
	'Get-TenantUsageSummaryStatistics',
	'Get-TenantVersion',
	'Get-User',
	'Get-UserByName',
	'Get-UserCollection',
	'Get-UserGroup',
	'Get-UserGroupByName',
	'Get-UserGroupCollection',
	'Get-UserGroupMembershipCollection',
	'Get-UserMembershipCollection',
	'Invoke-UserLogout',
	'New-Agent',
	'New-Alarm',
	'New-Application',
	'New-ApplicationBinary',
	'New-AuditRecord',
	'New-Binary',
	'New-BulkOperation',
	'New-Device',
	'New-DeviceGroup',
	'New-Event',
	'New-EventBinary',
	'New-ExternalId',
	'New-ManagedObject',
	'New-Measurement',
	'New-MicroserviceBinary',
	'New-Operation',
	'New-RetentionRule',
	'New-Tenant',
	'New-TenantOption',
	'New-User',
	'New-UserGroup',
	'Register-Device',
	'Remove-Agent',
	'Remove-AlarmCollection',
	'Remove-Application',
	'Remove-ApplicationBinary',
	'Remove-AssetFromGroup',
	'Remove-Binary',
	'Remove-BulkOperation',
	'Remove-ChildAddition',
	'Remove-ChildDeviceFromDevice',
	'Remove-Device',
	'Remove-DeviceFromGroup',
	'Remove-DeviceGroup',
	'Remove-DeviceRequest',
	'Remove-Event',
	'Remove-EventBinary',
	'Remove-EventCollection',
	'Remove-ExternalId',
	'Remove-ManagedObject',
	'Remove-Measurement',
	'Remove-MeasurementCollection',
	'Remove-Microservice',
	'Remove-OperationCollection',
	'Remove-RetentionRule',
	'Remove-RoleFromGroup',
	'Remove-RoleFromUser',
	'Remove-Tenant',
	'Remove-TenantOption',
	'Remove-User',
	'Remove-UserFromGroup',
	'Remove-UserGroup',
	'Request-DeviceCredentials',
	'Reset-UserPassword',
	'Set-DeviceRequiredAvailability',
	'Update-Agent',
	'Update-Alarm',
	'Update-AlarmCollection',
	'Update-Application',
	'Update-Binary',
	'Update-BulkOperation',
	'Update-CurrentApplication',
	'Update-CurrentUser',
	'Update-DataBrokerConnector',
	'Update-Device',
	'Update-DeviceGroup',
	'Update-Event',
	'Update-EventBinary',
	'Update-ManagedObject',
	'Update-Microservice',
	'Update-Operation',
	'Update-RetentionRule',
	'Update-Tenant',
	'Update-TenantOption',
	'Update-TenantOptionBulk',
	'Update-TenantOptionEditable',
	'Update-User',
	'Update-UserGroup',
	'Add-ClientResponseType',
	'Add-PowershellType',
	'Clear-Session',
	'ConvertFrom-Base64String',
	'ConvertFrom-ClientOutput',
	'ConvertFrom-JsonStream',
	'ConvertTo-Base64String',
	'ConvertTo-JsonArgument',
	'ConvertTo-NestedJson',
	'Expand-Application',
	'Expand-Device',
	'Expand-DeviceGroup',
	'Expand-Id',
	'Expand-Microservice',
	'Expand-PaginationObject',
	'Expand-Source',
	'Expand-Tenant',
	'Expand-User',
	'Find-ManagedObjectCollection',
	'Format-Date',
	'Get-AgentCollection',
	'Get-AssetParent',
	'Get-C8ySessionProperty',
	'Get-ClientBinary',
	'Get-ClientBinaryVersion',
	'Get-ClientCommonParameters',
	'Get-ClientSetting',
	'Get-CurrentTenantApplicationCollection',
	'Get-DeviceBootstrapCredential',
	'Get-DeviceCollection',
	'Get-DeviceGroupCollection',
	'Get-DeviceParent',
	'Get-ServiceUser',
	'Get-Session',
	'Get-SessionCollection',
	'Get-SessionHomePath',
	'Group-ClientRequests',
	'Install-ClientBinary',
	'Invoke-ClientIterator',
	'Invoke-ClientLogin',
	'Invoke-ClientRequest',
	'Invoke-NativeCumulocityRequest',
	'Invoke-Template',
	'New-HostedApplication',
	'New-Microservice',
	'New-RandomPassword',
	'New-RandomString',
	'New-ServiceUser',
	'New-Session',
	'New-TestAgent',
	'New-TestAlarm',
	'New-TestDevice',
	'New-TestDeviceGroup',
	'New-TestEvent',
	'New-TestFile',
	'New-TestMeasurement',
	'New-TestOperation',
	'New-TestUser',
	'New-TestUserGroup',
	'Open-Website',
	'Register-Alias',
	'Register-ClientArgumentCompleter',
	'Set-ClientConsoleSetting',
	'Set-Session',
	'Test-Json',
	'Unregister-Alias',
	'Wait-Operation',
	'Watch-Alarm',
	'Watch-Event',
	'Watch-ManagedObject',
	'Watch-Measurement',
	'Watch-Notification',
	'Watch-NotificationChannel',
	'Watch-Operation')
# VariablesToExport = '*'

# Aliases to export from this module, for best performance, do not use wildcards and do not delete the entry, use an empty array if there are no aliases to export.
AliasesToExport = '*'

# Private data to pass to the module specified in RootModule/ModuleToProcess. This may also contain a PSData hashtable with additional module metadata used by PowerShell.
PrivateData = @{
    PSData = @{
        #Prerelease of this module
        #Prerelease = '-alpha01'

        # Tags applied to this module. These help with module discovery in online galleries.
        Tags = @(
            'Cumulocity',
            'IoT',
            'REST-API',
            'Windows',
            'Linux',
            'MacOS',
            'PSEdition_Desktop',
            'PSEdition_Core'
        )

        # A URL to the license for this module.
        # LicenseUri = ''

        # A URL to the main website for this project.
        ProjectUri = 'https://reubenmiller.github.io/go-c8y-cli'

        # A URL to an icon representing this module.
        # IconUri = ''

        # ReleaseNotes of this module
        # ReleaseNotes = ''

        # External dependent modules of this module
        # ExternalModuleDependencies = ''

    } # End of PSData hashtable

} # End of PrivateData hashtable

# HelpInfo URI of this module
# HelpInfoURI = ''

# Default prefix for commands exported from this module. Override the default prefix using Import-Module -Prefix.
DefaultCommandPrefix = ''

}
