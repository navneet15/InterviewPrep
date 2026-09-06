# Rate Limiter

## Problem Statement
Design and implement a rate limiter that throttles incoming requests per
client per endpoint, according to pre-configured rules, supporting multiple
rate-limiting algorithms.

## Requirements
1. Each request for any endpoint must be rate-limited as per pre-defined
   sets of rules based on `clientId`.
2. Configuration is loaded on application bootup. No dynamic updates are
   required.
3. Token Bucket and Sliding Window algorithms are required.
4. System receives requests with:
   - `clientId`: string
   - `endpoint`: string
5. Return a structured response for every request:
   - `allowed`: bool
   - `remaining`: int
   - `retryAfterMs`: int64 | null
6. If an endpoint has no configured rule, fall back to a default rate limit.

## Entities / Classes

1. **RateLimiter**
   Main coordinator of the system.
   - Receives `clientId` and `endpoint`.
   - Finds the correct `RateLimitRule`.
   - Finds or creates the runtime limiter for that client + endpoint.
   - Delegates the rate-limit decision to the appropriate strategy.
   - Returns whether the request should be allowed or rejected.

2. **RateLimitStrategy**
   Represents the common behaviour shared by all rate-limiting algorithms —
   "decide whether the current request should be allowed." Implemented by
   `TokenBucketLimiter` and `SlidingWindowLimiter`.

3. **RateLimitRule**
   Represents one configured rate-limiting policy. Groups an endpoint, the
   selected algorithm, and algorithm-specific configuration.
   Example: `/payments` → Algorithm: TokenBucket, Capacity: 10, RefillRate: 2/sec.

4. **SlidingWindowLimiter**
   Owns the configuration, runtime state, and business logic for the Sliding
   Window algorithm. Runtime state may contain request timestamps within the
   configured time window.

5. **TokenBucketLimiter**
   Owns the configuration, runtime state, and business logic for the Token
   Bucket algorithm. Runtime state may contain the current token count and
   the last refill timestamp.

6. **RateLimitKey**
   Runtime limiter state must be maintained independently per
   `clientId + endpoint`. `RateLimitKey` represents this combination as a
   single meaningful value, e.g. `{ClientID: "client-123", Endpoint: "/payments"}`.
