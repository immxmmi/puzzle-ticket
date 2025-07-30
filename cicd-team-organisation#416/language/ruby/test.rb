require 'sentry-ruby'

require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'

OpenTelemetry::SDK.configure do |c|
  c.service_name = 'ruby-glitchtip'
  c.use_all
end

tracer = OpenTelemetry.tracer_provider.tracer('glitchtip-ruby-tracer')

tracer.in_span('test-span') do |span|
  puts "Tracing inside test-span"
end

Sentry.init do |config|
  config.dsn = 'https://69d0738116074e24a0f6a0727834737e@glitchtip-stg.puzzle.ch/10'
  config.breadcrumbs_logger = [:active_support_logger, :http_logger]
  config.environment = 'development'
  config.debug = true
end

begin
  raise "Test error for GlitchTip Ruby"
rescue => e
  Sentry.capture_exception(e)
  raise
end