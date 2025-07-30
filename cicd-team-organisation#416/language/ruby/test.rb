require 'sentry-ruby'

Sentry.init do |config|
  config.dsn = 'https://69d0738116074e24a0f6a0727834737e@glitchtip-stg.puzzle.ch/10'
  config.breadcrumbs_logger = [:active_support_logger, :http_logger]
end

raise "Test error for GlitchTip Ruby"