﻿Function New-C8yPowershellArguments {
    [cmdletbinding()]
    Param(
        [Parameter(
            Mandatory = $true,
            Position = 0
        )]
        [string] $Name,

        [Parameter(
            Mandatory = $true,
            Position = 1
        )]
        [string] $Type,

        [string] $Required,

        [string] $OptionName,

        [string] $Description,

        [string] $Default,

        [switch] $ReadFromPipeline
    )

    # Format variable name
    $NameLocalVariable = $Name[0].ToString().ToUpperInvariant() + $Name.Substring(1)

    $ParameterDefinition = New-Object System.Collections.ArrayList

    if ($Required -match "true|yes") {
        $null = $ParameterDefinition.Add("Mandatory = `$true")
        $Description = "${Description} (required)"
    }

    # TODO: Do we need to add Position = x? to the ParameterDefinition

    # Add alias
    if ($UseOption) {
        $null = $ParameterDefinition.Add("Alias = `"$OptionName`"")
    }

    # Add Piped argument
    if ($Type -match "(source|id)" -or $ReadFromPipeline) {
        $null = $ParameterDefinition.Add("ValueFromPipeline=`$true")
        $null = $ParameterDefinition.Add("ValueFromPipelineByPropertyName=`$true")
    }

    # Type Definition
    $DataType = switch ($Type) {
        "[]agent" { "object[]"; break }
        "[]device" { "object[]"; break }
        "[]devicegroup" { "object[]"; break }
        "[]id" { "object[]"; break }
        "[]smartgroup" { "object[]"; break }
        "[]role" { "object[]"; break }
        "[]roleself" { "object[]"; break }
        "[]string" { "string[]"; break }
        "[]stringcsv" { "string[]"; break }
        "[]tenant" { "object[]"; break }
        "[]user" { "object[]"; break }
        "[]usergroup" { "object[]"; break }
        "[]userself" { "object[]"; break }
        "application" { "object[]"; break }
        "boolean" { "switch"; break }
        "optional_fragment" { "switch"; break }
        "datefrom" { "string"; break }
        "datetime" { "string"; break }
        "dateto" { "string"; break }
        "directory" { "string"; break }
        "file" { "string"; break }
        "float" { "float"; break }
        "fileContents" { "string"; break }
        "attachment" { "string"; break }
        "id" { "object[]"; break }
        "integer" { "long"; break }
        "json" { "object"; break }
        "json_custom" { "object"; break }
        "microservice" { "object[]"; break }
        "microserviceinstance" { "string"; break }
        "set" { "object[]"; break }
        "source" { "object"; break }
        "string" { "string"; break }
        "[]devicerequest" { "object[]"; break }
        "strings" { "string"; break }
        "tenant" { "object"; break }
        default {
            Write-Error "Unsupported Type. $_"
        }
    }

    New-Object psobject -Property @{
        Name = $NameLocalVariable
        Type = $DataType
        Definition = $ParameterDefinition
        Description = "$Description"

    }
}
