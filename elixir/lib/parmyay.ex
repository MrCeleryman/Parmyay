defmodule Parmyay do
	use Application
	require Logger

	def start(_type, _args) do
		port = Application.get_env(:parmyay, :cowboy_port)
		
		children = [
			Plug.Adapters.Cowboy.child_spec(:http, Parmyay.Routers, [], port: port)
		]
		Logger.info "Started parmyay"
		Supervisor.start_link(children, strategy: :one_for_one)
	end

end
