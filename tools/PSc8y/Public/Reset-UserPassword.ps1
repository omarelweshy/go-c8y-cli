﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Reset-UserPassword {
<#
.SYNOPSIS
Reset user password

.DESCRIPTION
The password can be reset either by issuing a password reset email (default), or be specifying a new password.

Note: In more recent Cumulocity IoT versions,  you can't set a fixed password for another user.


.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/users_resetUserPassword

.EXAMPLE
PS> Reset-UserPassword -Id $User.id -Dry

Resets a user's password by sending a reset email to the user

.EXAMPLE
PS> Reset-UserPassword -Id $User.id -NewPassword (New-RandomPassword)

Resets a user's password by generating a new password


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # User id (required)
        [Parameter(Mandatory = $true,
                   ValueFromPipeline=$true,
                   ValueFromPipelineByPropertyName=$true)]
        [object[]]
        $Id,

        # New user password. Min: 6, max: 32 characters. Only Latin1 chars allowed
        [Parameter()]
        [string]
        $NewPassword,

        # Tenant
        [Parameter()]
        [object]
        $Tenant
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Update", "Template"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "users resetUserPassword"
        $ClientOptions = Get-ClientOutputOption $PSBoundParameters
        $TypeOptions = @{
            Type = "application/vnd.com.nsn.cumulocity.user+json"
            ItemType = ""
            BoundParameters = $PSBoundParameters
        }
    }

    Process {

        if ($ClientOptions.ConvertToPS) {
            $Id `
            | Group-ClientRequests `
            | c8y users resetUserPassword $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            $Id `
            | Group-ClientRequests `
            | c8y users resetUserPassword $c8yargs
        }
        
    }

    End {}
}
