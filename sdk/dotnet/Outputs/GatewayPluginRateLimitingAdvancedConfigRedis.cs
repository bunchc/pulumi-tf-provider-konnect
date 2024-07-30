// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Outputs
{

    [OutputType]
    public sealed class GatewayPluginRateLimitingAdvancedConfigRedis
    {
        /// <summary>
        /// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Cluster. Each string element must be a hostname. The minimum length of the array is 1 element.
        /// </summary>
        public readonly ImmutableArray<string> ClusterAddresses;
        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        public readonly int? ConnectTimeout;
        /// <summary>
        /// Database to use for the Redis connection when using the `redis` strategy
        /// </summary>
        public readonly int? Database;
        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
        /// </summary>
        public readonly int? KeepaliveBacklog;
        /// <summary>
        /// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
        /// </summary>
        public readonly int? KeepalivePoolSize;
        /// <summary>
        /// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// An integer representing a port number between 0 and 65535, inclusive.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        public readonly int? ReadTimeout;
        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        public readonly int? SendTimeout;
        /// <summary>
        /// Sentinel addresses to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel. Each string element must be a hostname. The minimum length of the array is 1 element.
        /// </summary>
        public readonly ImmutableArray<string> SentinelAddresses;
        /// <summary>
        /// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
        /// </summary>
        public readonly string? SentinelMaster;
        /// <summary>
        /// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
        /// </summary>
        public readonly string? SentinelPassword;
        /// <summary>
        /// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel. must be one of ["master", "slave", "any"]
        /// </summary>
        public readonly string? SentinelRole;
        /// <summary>
        /// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
        /// </summary>
        public readonly string? SentinelUsername;
        /// <summary>
        /// A string representing an SNI (server name indication) value for TLS.
        /// </summary>
        public readonly string? ServerName;
        /// <summary>
        /// If set to true, uses SSL to connect to Redis.
        /// </summary>
        public readonly bool? Ssl;
        /// <summary>
        /// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
        /// </summary>
        public readonly bool? SslVerify;
        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private GatewayPluginRateLimitingAdvancedConfigRedis(
            ImmutableArray<string> clusterAddresses,

            int? connectTimeout,

            int? database,

            string? host,

            int? keepaliveBacklog,

            int? keepalivePoolSize,

            string? password,

            int? port,

            int? readTimeout,

            int? sendTimeout,

            ImmutableArray<string> sentinelAddresses,

            string? sentinelMaster,

            string? sentinelPassword,

            string? sentinelRole,

            string? sentinelUsername,

            string? serverName,

            bool? ssl,

            bool? sslVerify,

            int? timeout,

            string? username)
        {
            ClusterAddresses = clusterAddresses;
            ConnectTimeout = connectTimeout;
            Database = database;
            Host = host;
            KeepaliveBacklog = keepaliveBacklog;
            KeepalivePoolSize = keepalivePoolSize;
            Password = password;
            Port = port;
            ReadTimeout = readTimeout;
            SendTimeout = sendTimeout;
            SentinelAddresses = sentinelAddresses;
            SentinelMaster = sentinelMaster;
            SentinelPassword = sentinelPassword;
            SentinelRole = sentinelRole;
            SentinelUsername = sentinelUsername;
            ServerName = serverName;
            Ssl = ssl;
            SslVerify = sslVerify;
            Timeout = timeout;
            Username = username;
        }
    }
}
