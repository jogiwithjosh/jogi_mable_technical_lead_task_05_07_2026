import { PassThrough } from 'node:stream';

import { renderToPipeableStream } from 'react-dom/server';

import { RemixServer } from '@remix-run/react';

import type { EntryContext } from '@remix-run/node';

export default function handleRequest(
  request: Request,
  responseStatusCode: number,
  responseHeaders: Headers,
  remixContext: EntryContext
) {
  return new Promise<Response>((resolve, reject) => {
    let didError = false;

    const { pipe } = renderToPipeableStream(
      <RemixServer
        context={remixContext}
        url={request.url}
      />,
      {
        onShellReady() {
          const body = new PassThrough();

          responseHeaders.set(
            'Content-Type',
            'text/html'
          );

          resolve(
            new Response(body as any, {
              headers: responseHeaders,
              status: didError
                ? 500
                : responseStatusCode
            })
          );

          pipe(body);
        },

        onShellError(error) {
          reject(error);
        },

        onError(error) {
          didError = true;
          console.error(error);
        }
      }
    );
  });
}