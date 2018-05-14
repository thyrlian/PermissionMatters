# PermissionMatters

[![Go Report Card](https://goreportcard.com/badge/github.com/thyrlian/PermissionMatters)](https://goreportcard.com/report/github.com/thyrlian/PermissionMatters)
[![Android开发技术周报](https://img.shields.io/badge/Android%E5%BC%80%E5%8F%91%E6%8A%80%E6%9C%AF%E5%91%A8%E6%8A%A5-%23179-yellowgreen.svg)](https://androidweekly.io/android-dev-weekly-issue-179/)

Check your Android application's permission changes

[Permission](https://developer.android.com/guide/topics/permissions/index.html) matters!  That's what the users care the most.  Users are always suspicious of why on earth an application needs any of the permissions, they hate every permission from the bottom of heart.  It's a good practice to check if there is any new permission added to your Android application by any means (e.g. updating 3rd party library could also bring some new permissions).  When it happens, the adoption of application's new version will be affected.

This script could help warn developers of any permission change by setting up a CI job.

## Precondition

* [Android SDK](https://developer.android.com/studio/index.html#command-tools) ([`apkanalyzer`](https://developer.android.com/studio/command-line/apkanalyzer.html)) is installed
* `ANDROID_HOME` environment variable is set (or at least pass in command: `export ANDROID_HOME=<PATH_TO_YOUR_ANDROID_SDK>`)

## Build

```console
scripts/build.sh

# => GOPATH: <YOUR_CLONED_PATH>/PermissionMatters/src
# => GOBIN:  <YOUR_CLONED_PATH>/PermissionMatters/bin

# => Binary file is generated to: <YOUR_CLONED_PATH>/PermissionMatters/bin/permissionguard
```

Or simply grab the built binary [here](https://github.com/thyrlian/PermissionMatters/blob/master/bin/permissionguard).

## Play

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

## Results

There are 4 possible cases (`++`, `--`, `++ & --`, `==`), and the result examples are like below:

* permission++
* => Fail (exit code 1)
```console
======================================================================
Failure!

4 new permission(s) added:
    android.permission.CAMERA
    android.permission.FLASHLIGHT
    android.permission.SEND_SMS
    com.me.app.myapp.permission.DEADLY_ACTIVITY
======================================================================
```

* permission++  &  permission--
* => Fail (exit code 1)
```console
======================================================================
Failure!

4 new permission(s) added:
    android.permission.WRITE_EXTERNAL_STORAGE
    com.sonyericsson.home.permission.BROADCAST_BADGE
    com.sec.android.provider.badge.permission.READ
    com.sec.android.provider.badge.permission.WRITE

2 old permission(s) removed:
    android.permission.CAMERA
    android.permission.FLASHLIGHT
======================================================================
```

* permission--
* => Warn (exit code 1)
* TODO: you need to update the snapshot (take a new one)
```console
======================================================================
Warning!

3 old permission(s) removed:
    android.permission.CAMERA
    android.permission.FLASHLIGHT
    com.me.app.myapp.permission.DEADLY_ACTIVITY

A new snapshot needs to be taken.
======================================================================
```

* permission==
* => Pass (exit code 0)
```console
======================================================================
No permission is changed.
======================================================================
```

## License

Copyright (c) 2018 Jing Li. See the [LICENSE](https://github.com/thyrlian/PermissionMatters/blob/master/LICENSE) file for license rights and limitations (MIT).
