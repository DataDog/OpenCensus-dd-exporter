// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.
package datadog

import (
	"testing"
	"time"

	"go.opencensus.io/stats/view"
)

func TestAddViewData(t *testing.T) {
	exporter := newExporter(Options{Namespace: "hello", Tags: []string{"test:optionalTag"}})

	view.RegisterExporter(exporter)
	reportPeriod := time.Millisecond
	view.SetReportingPeriod(reportPeriod)

	expected := &view.Data{
		View: newView(view.Count()),
	}
	exporter.collector.addViewData(expected)
	actual := exporter.collector.viewData["hello.fooCount"]

	if actual != expected {
		t.Errorf("Expected: %v, Got: %v", expected, actual)
	}
}
