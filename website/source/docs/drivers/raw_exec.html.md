---
layout: "docs"
page_title: "Drivers: Raw Exec"
sidebar_current: "docs-drivers-raw-exec"
description: |-
  The Raw Exec task driver simply fork/execs and provides no isolation.
---

# Raw Fork/Exec Driver

Name: `raw_exec`

The `raw_exec` driver is used to execute a command for a task without any
isolation. Further, the task is started as the same user as the Nomad process.
As such, it should be used with extreme care and is disabled by default.

## Task Configuration

The `raw_exec` driver supports the following configuration in the job spec:

* `command` - The command to execute. Must be provided.
* `artifact_source` – Source location of an executable artifact. Must be accessible
from the Nomad client
* `args` - The argument list to the command, space seperated. Optional.

## Client Requirements

The `raw_exec` driver can run on all supported operating systems. It is however
disabled by default. In order to be enabled, the Nomad client configuration must
explicitly enable the `raw_exec` driver in the
[options](../agent/config.html#options) field:

```
options = {
    driver.raw_exec.enable = "1"
}
```

You must specify either a `command` or a `artifact_source` to be executed. Any
`command` is assumed to be present on the running client.

## Client Attributes

The `raw_exec` driver will set the following client attributes:

* `driver.raw_exec` - This will be set to "1", indicating the
  driver is available.

## Resource Isolation

The `raw_exec` driver provides no isolation.
