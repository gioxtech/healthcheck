# Healthcheck

This is a really simple tool to perform a healthcheck. For now the only implemented protocol is `http`, but it's possible to extend to others protocols.

If you're asking yourself why do you need that? If you say yes for any of these question:

- You have a really tiny docker image (from scratch or alpine) and you don't want to install curl for example.
- You don't want to develop your own health command line.
- You want a easy way to add healthcheck to your services.

To make your life easy we already have staticaly pre-compiled releases, to start to use add the following commands to your `Dockerfile`:

```dockerfile
ENV HEALTHCHECKER_VERSION 1.0.0
ENV HEALTHCHECKER_URL https://github.com/gioxtech/healthcheck/releases/download/v${HEALTHCHECKER_VERSION}/healthcheck-${HEALTHCHECKER_VERSION}
RUN wget ${HEALTHCHECKER_URL} -O /usr/bin/healthcheck

HEALTHCHECK --start-period=15s --interval=15s --timeout=1s --retries=6 CMD healthcheck -host-addr http://localhost/health
```
