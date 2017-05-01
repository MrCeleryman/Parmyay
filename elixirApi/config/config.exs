# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# Configures the endpoint
config :parmyay, Parmyay.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "DhOznzMakv/fAaj1E1hNMhi8m305+8aL8hNdGuP1cp5sA2E1cEsO8vzXztTgKAHj",
  render_errors: [view: Parmyay.ErrorView, accepts: ~w(html json)],
  pubsub: [name: Parmyay.PubSub,
           adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
