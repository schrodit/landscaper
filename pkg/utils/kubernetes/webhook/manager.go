// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"context"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	manager2 "sigs.k8s.io/controller-runtime/pkg/manager"
)

type Webhook interface {
	Name() string
	Handler() http.Handler
	WebhookConfiguration(namespace, serviceName string, servicePort int) (runtime.Object, error)
}

type Manager interface {
	Register(webhook Webhook) error
}

// manager implements the Manager interface for webhooks
type manager struct {
	mgr manager2.Manager
	webhooks map[string]Webhook
}

func New(mgr manager2.Manager) Manager {
	return &manager{
		mgr: mgr,
		webhooks: map[string]Webhook{},
	}
}

func (m *manager) Register(webhook Webhook) error {
	m.webhooks[webhook.Name()] = webhook
	return nil
}

func (m *manager) Complete(ctx context.Context, namespace, serviceName string, servicePort int) error {
	for _, webhook := range m.webhooks {
		m.mgr.GetWebhookServer().Register("/validate-" + webhook.Name(), webhook.Handler())
	}

	return nil
}
