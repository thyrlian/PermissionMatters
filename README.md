# PermissionMatters

Check your Android application's permission changes

[Permission](https://developer.android.com/guide/topics/permissions/index.html) matters!  That's what the users care the most.  Users are always suspicious of why on earth an application needs any of the permissions, they hate every permission from the bottom of heart.  It's a good practice to check if there is any new permission added to your Android application by any means (e.g. updating 3rd party library could also bring some new permissions).  When it happens, the adoption of application's new version will be affected.

This script could help warn developers of any permission change by setting up a CI job.

## Precondition

* [Android SDK](https://developer.android.com/studio/index.html#command-tools) ([`apkanalyzer`](https://developer.android.com/studio/command-line/apkanalyzer.html))
