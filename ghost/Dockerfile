ARG VERSION=3.32.0
FROM ghost:${VERSION}-alpine
ARG VERSION
COPY src/ghost.min* ./versions/${VERSION}/core/built/assets/
# ARG VERSION
COPY src/vendor.min* ./versions/${VERSION}/core/built/assets/
# ARG VERSION
COPY src/content-cta.hbs ./versions/${VERSION}/core/frontend/helpers/tpl/
COPY src/maintenance.html ./versions/${VERSION}/core/server/views/
COPY src/user.js ./versions/${VERSION}/core/server/models/
COPY src/GhostMailer.js ./versions/${VERSION}/core/server/services/mail/
COPY src/api.js ./versions/${VERSION}/core/server/services/members/
COPY src/emails/* ./versions/${VERSION}/core/server/services/members/emails/
