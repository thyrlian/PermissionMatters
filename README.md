# PermissionMatters

[![Go Report Card](https://goreportcard.com/badge/github.com/thyrlian/PermissionMatters)](https://goreportcard.com/report/github.com/thyrlian/PermissionMatters)

Check your Android application's permission changes

[Permission](https://developer.android.com/guide/topics/permissions/index.html) matters!  That's what the users care the most.  Users are always suspicious of why on earth an application needs any of the permissions, they hate every permission from the bottom of heart.  It's a good practice to check if there is any new permission added to your Android application by any means (e.g. updating 3rd party library could also bring some new permissions).  When it happens, the adoption of application's new version will be affected.

This script could help warn developers of any permission change by setting up a CI job.

## Precondition

* [Android SDK](https://developer.android.com/studio/index.html#command-tools) ([`apkanalyzer`](https://developer.android.com/studio/command-line/apkanalyzer.html)) is installed
* `ANDROID_HOME` environment variable is set (or at least pass in command: `export ANDROID_HOME=<PATH_TO_YOUR_ANDROID_SDK>`)

## How To Use

```console
permissionguard <take|scan> -apk <apk> [-snapshot <your_snapshot_file_of_permissions>]

Subject   Description
-------   -----------
take      Take snapshot of the given APK's permissions (for the first time or after any permission change)
scan      Scan the given APK file, compare its permissions with the snapshot

Option    Description
------    -----------
apk       The APK file to analyze
snapshot  The permission snapshot file (default "./permissions.json")
```

## License

Copyright (c) 2018 Jing Li. See the [LICENSE](https://github.com/thyrlian/PermissionMatters/blob/master/LICENSE) file for license rights and limitations (MIT).
