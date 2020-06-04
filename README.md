<div align="center">
	<img width="500" src=".github/logo.svg" alt="pinpt-logo">
</div>

<p align="center" color="#6a737d">
	<strong>This repo contains Golang utilities for working with Atlassian Document Format content</strong>
</p>


## Overview

The [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure/) (ADF) represents rich text stored in Atlassian products. For example, in Jira Cloud platform, the text in issue comments and in textarea custom fields is stored as ADF.

This library provides a set of utilities in Golang for dealing with ADF content.


## Install

```
go get -u github.com/pinpt/adf
```

## Usage

import github.com/pinpt/adf

```golang
buf, err := adf.GenerateHTMLFromADF([]byte(`{
  "version": 1,
  "type": "doc",
  "content": [
    {
      "type": "paragraph",
      "content": [
        {
          "type": "text",
          "text": "Hello world"
        }
      ]
    }
  ]
}`))
```

## License

Licensed under the MIT License.
