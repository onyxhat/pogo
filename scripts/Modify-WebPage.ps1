Param (
    [string]$Uri = "www.yahoo.com",
    [string]$Find,
    [string]$Replace
)

$req = Invoke-WebRequest -Uri $Uri

$req.Content -replace $Find,$Replace
