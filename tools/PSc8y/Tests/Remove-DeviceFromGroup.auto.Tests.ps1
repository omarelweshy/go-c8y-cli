﻿. $PSScriptRoot/imports.ps1

Describe -Name "Remove-DeviceFromGroup" {
    BeforeEach {
        $Group = PSc8y\New-TestDeviceGroup
        $ChildDevice = PSc8y\New-TestDevice
        PSc8y\Add-AssetToGroup -Group $Group.id -Child $ChildDevice.id

    }

    It "Unassign a child device from its parent asset" {
        $Response = PSc8y\Remove-DeviceFromGroup -Group $Group.id -Child $ChildDevice.id
        $LASTEXITCODE | Should -Be 0
    }


    AfterEach {
        PSc8y\Remove-ManagedObject -Id $ChildDevice.id
        PSc8y\Remove-ManagedObject -Id $Group.id

    }
}

