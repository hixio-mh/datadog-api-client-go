# \UsageMeteringApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageFargate**](UsageMeteringApi.md#GetUsageFargate) | **Get** /api/v1/usage/fargate | Get hourly usage for Fargate
[**GetUsageHosts**](UsageMeteringApi.md#GetUsageHosts) | **Get** /api/v1/usage/hosts | Get hourly usage for hosts and containers
[**GetUsageLogs**](UsageMeteringApi.md#GetUsageLogs) | **Get** /api/v1/usage/logs | Get hourly usage for Logs
[**GetUsageRumSessions**](UsageMeteringApi.md#GetUsageRumSessions) | **Get** /api/v1/usage/rum_sessions | Get hourly usage for RUM Sessions
[**GetUsageSummary**](UsageMeteringApi.md#GetUsageSummary) | **Get** /api/v1/usage/summary | Get usage across your multi-org account
[**GetUsageSynthetics**](UsageMeteringApi.md#GetUsageSynthetics) | **Get** /api/v1/usage/synthetics | Get hourly usage for Synthetics
[**GetUsageTimeseries**](UsageMeteringApi.md#GetUsageTimeseries) | **Get** /api/v1/usage/timeseries | Get hourly usage for custom metrics
[**GetUsageTopAvgMetrics**](UsageMeteringApi.md#GetUsageTopAvgMetrics) | **Get** /api/v1/usage/top_avg_metrics | Get top 500 custom metrics by hourly average
[**GetUsageTrace**](UsageMeteringApi.md#GetUsageTrace) | **Get** /api/v1/usage/traces | Get hourly usage for Trace Search



## GetUsageFargate

> UsageFargateResponse GetUsageFargate(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Fargate



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageFargateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageFargateResponse**](UsageFargateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageHosts

> UsageHostsResponse GetUsageHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for hosts and containers



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageHostsResponse**](UsageHostsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogs

> UsageLogsResponse GetUsageLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Logs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageLogsResponse**](UsageLogsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageRumSessions

> UsageRumSessionsResponse GetUsageRumSessions(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for RUM Sessions



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRumSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageRumSessionsResponse**](UsageRumSessionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> UsageSummaryResponse GetUsageSummary(ctx).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()

Get usage across your multi-org account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning in this month. Maximum of 15 months ago. | 
 **endMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage ending this month. | 
 **includeOrgDetails** | **bool** | Include usage summaries for each sub-org. | 

### Return type

[**UsageSummaryResponse**](UsageSummaryResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSynthetics

> UsageSyntheticsResponse GetUsageSynthetics(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Synthetics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSyntheticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageSyntheticsResponse**](UsageSyntheticsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTimeseries

> UsageTimeseriesResponse GetUsageTimeseries(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for custom metrics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageTimeseriesResponse**](UsageTimeseriesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTopAvgMetrics

> UsageTopAvgMetricsResponse GetUsageTopAvgMetrics(ctx).Month(month).Names(names).Execute()

Get top 500 custom metrics by hourly average



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTopAvgMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning at this hour. | 
 **names** | [**[]string**](string.md) | Comma-separated list of metric names. | 

### Return type

[**UsageTopAvgMetricsResponse**](UsageTopAvgMetricsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTrace

> UsageTraceResponse GetUsageTrace(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Trace Search



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending BEFORE this hour | 

### Return type

[**UsageTraceResponse**](UsageTraceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
