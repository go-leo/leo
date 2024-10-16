# Bulkhead
it can prevent issues in one service area from affecting the entire service.

The idea behind Bulkhead is to set limits on the number of concurrent calls we make to remote services. We treat calls to different remote services as separate, isolated pools and set limits on the number of calls that can be made concurrently.