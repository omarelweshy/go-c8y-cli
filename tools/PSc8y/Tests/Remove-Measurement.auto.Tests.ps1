﻿. $PSScriptRoot/imports.ps1

Describe -Name "Remove-Measurement" {
    BeforeEach {
        $Measurement = New-TestDevice | New-Measurement -Template "test.measurement.jsonnet"

    }

    It -Skip "Delete measurement" {
        $Response = PSc8y\Remove-Measurement -id $Measurement.id
        $LASTEXITCODE | Should -Be 0
    }


    AfterEach {
        PSc8y\Remove-ManagedObject -Id $Measurement.source.id

    }
}

