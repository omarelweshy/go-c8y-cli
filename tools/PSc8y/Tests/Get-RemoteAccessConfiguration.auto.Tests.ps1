﻿. $PSScriptRoot/imports.ps1

Describe -Name "Get-RemoteAccessConfiguration" {
    BeforeEach {

    }

    It -Skip "Get existing remote access configuration" {
        $Response = PSc8y\Get-RemoteAccessConfiguration -Device device01 -Id 1
        $LASTEXITCODE | Should -Be 0
    }


    AfterEach {

    }
}

