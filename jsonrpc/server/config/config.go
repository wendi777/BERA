// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package config

import "time"

type (
	// `Server` defines the configuration for the JSON-RPC server.
	Server struct {
		// `API` defines a list of JSON-RPC namespaces to be enabled.
		EnableAPIs []string `mapstructure:"api"`

		// `Address` defines the HTTP server to listen on.
		Address string `mapstructure:"address"`

		// `WsAddress` defines the WebSocket server to listen on.
		WSAddress string `mapstructure:"ws-address"`

		// `MetricsAddress` defines the metrics server to listen on.
		MetricsAddress string `mapstructure:"metrics-address"`

		// `HTTPReadHeaderTimeout` is the read timeout of http json-rpc server.
		HTTPReadHeaderTimeout time.Duration `mapstructure:"http-read-header-timeout"`

		// `HTTPReadTimeout` is the read timeout of http json-rpc server.
		HTTPReadTimeout time.Duration `mapstructure:"http-read-timeout"`

		// `HTTPWriteTimeout` is the write timeout of http json-rpc server.
		HTTPWriteTimeout time.Duration `mapstructure:"http-write-timeout"`

		// HTTPIdleTimeout is the idle timeout of http json-rpc server.
		HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`

		// `HTTPBaseRoute` defines the base path for the JSON-RPC server.
		BaseRoute string `mapstructure:"base-path"`

		// `TLSConfig` defines the TLS configuration for the JSON-RPC server.
		TLSConfig *TLSConfig `mapstructure:"tls-config"`
	}

	// `TLSConfig` defines a certificate and matching private key for the server.
	TLSConfig struct {
		// `CertPath` the file path for the certificate .pem file
		CertPath string `mapstructure:"cert-path"`

		// KeyPath the file path for the key .pem file
		KeyPath string `toml:"key-path"`
	}
)
