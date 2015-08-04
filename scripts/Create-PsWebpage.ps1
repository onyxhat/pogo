Param (
    [string]$FirstName,
    [string]$LastName
)

###Templates
$Inline = @"
<html>
    <h1>$(Split-Path $MyInvocation.MyCommand.Definition -Leaf)</h1>

    <form method="GET">
        <fieldset>
            <legend>Parameters</legend>
            $($MyInvocation.MyCommand.Parameters.Keys | % { "<label>$_ <input type='text' name='-$_'></label><br>`r`n" })
        </fieldset>
        <input type="submit" value="Submit">
    </form>
</html>
"@

$Result = @"
    <h1>Hello, $FirstName $LastName</h1>

    <p>Thanks for testing this $FirstName</p>
"@

###Runtime
if ($($PSBoundParameters.Keys)) {
    Write-Output $Result
} else {
    Write-Output $Inline
}
