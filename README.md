# redhat-secdata
Demonstrates how to use [Clair Core](https://github.com/quay/claircore) to aggregate and search Red Hat Security Data. 

## Known Limitations
Unable to provide accurate data for RHEL EUS streams such as RHEL 8.1. A RHEL 8.1 EUS system "should" be finding "RHEL 8.1 EUS" definitions and "RHEL 8.0 base" definitions and "RHEL 8.1 base" definitions, and the CPEs that Red Hat publishes are not specific enough to determine those last two; all "RHEL 8 base" errata are simply released with a "RHEL 8" CPE. So you cannot know, based on current CPEs alone, which of the "RHEL 8" CPEs you are supposed to compare against a "RHEL 8.1 EUS" system.
