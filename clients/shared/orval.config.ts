// packages/shared/orval.config.ts
import { defineConfig } from 'orval';

export default defineConfig({
    api: {
        input: { target: '../../backend/api/openapi.yaml' },
        output: {
            mode: 'tags-split',
            target: 'src/api/generated/endpoints',
            schemas: 'src/api/generated/model',
            client: 'react-query',
            httpClient: 'fetch',
            mock: true,
            clean: true,
            override: {
                // mutator: { path: 'src/api/mutator.ts', name: 'customFetch' },
                fetch: { includeHttpResponseReturnType: false },
                query: { useQuery: true, signal: true },
                operations: {
                    listOrders: {
                        query: { useInfinite: true, useInfiniteQueryParam: 'cursor' },
                    },
                    listMessages: {
                        query: { useInfinite: true, useInfiniteQueryParam: 'cursor' },
                    },
                },
            },
        },
    },
    // опционально: zod-схемы из той же спеки для форм
    // apiZod: {
    //     input: { target: '../../api/openapi.yaml' },
    //     output: { mode: 'tags-split', target: 'src/api/generated/zod', client: 'zod' },
    // },
});