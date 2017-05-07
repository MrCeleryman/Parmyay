defmodule Parmyay.Routers do
  	use Plug.Router

	if Mix.env == :dev do
    	use Plug.Debugger
  	end

	use Plug.ErrorHandler
	
	plug :match
	plug :dispatch

	match "/users" do
		{conn, status, body} = conn |> handle_request
		send_resp(conn, status, body)
	end

	def handle_request(%Plug.Conn{method: method} = conn) when method == "GET" do
		{conn, 200, "wew lad"}
	end

	# forward "/users", to: Parmyay.Routers.Users

	forward "/venues", to: Parmyay.Routers.Venues

	forward "/reviews", to: Parmyay.Routers.Reviews

	forward "/achievements", to: Parmyay.Routers.Achievements

	match _, do: send_resp(conn, 404, "Oops!")

	defp handle_errors(conn, %{kind: _kind, reason: _reason, stack: _stack}) do
    	send_resp(conn, conn.status, "Something went wrong")
  	end
end
