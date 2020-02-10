// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// The examples/query folder contains integration tests for all SQL statements, data types supported and unsupported by AWS Athena.
// How to prepare integration end-to-end test?
//
// 1. Prerequisites - AWS Credentials & S3 Query Result Bucket.
//
// To be able to query AWS Athena, you need to have an AWS account at Amazon AWS's website. To give it a shot,
// a free tier account is enough. You also need to have a pair of AWS access key ID and secret access key.
// You can get it from AWS Security Credentials section of Identity and Access Management (IAM).
// If you don't have one, please create it.
//
//
// In addition to AWS credentials, you also need an s3 bucket to store query result.
// Just go to AWS S3 web console page to create one. In the examples below,
// the s3 bucket I use is s3://henrywuqueryresults/.
//
// In most cases, you need the following 4 prerequisites :
//
//     S3 Output bucket
//     access key ID
//     secret access key
//     AWS region
//
// For more details on athenasql's support on AWS credentials & S3 query result bucket,
// please refer to README section Support Multiple AWS Authorization Methods.
//
// 2. Installation athenasql.
//
//     go get github.com/uber/athenasql
//
// 3. Integration Test.
//
// To Build it:
//
//     $cd $GOPATH/src/github.com/uber/athenasql
//     $go build examples/query/dml_select_simple.go
//
//
// Run it and you can see output like:
//
//    $./dml_select_simple
//    https://www.example.com/jobs/433
//
package main
