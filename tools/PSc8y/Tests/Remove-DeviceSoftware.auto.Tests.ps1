﻿. $PSScriptRoot/imports.ps1

Describe -Name "Remove-DeviceSoftware" {
    BeforeEach {

    }

    It -Skip "Remove software" {
        $Response = PSc8y\Remove-DeviceSoftware -Id 12345 -Name ntp -Version 1.0.0
        $LASTEXITCODE | Should -Be 0
        $Response | Should -Not -BeNullOrEmpty
    }


    AfterEach {

    }
}

