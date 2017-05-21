defmodule Parmyay.Router do
  	use Plug.Router

	if Mix.env == :dev do
    	use Plug.Debugger
  	end

	use Plug.ErrorHandler

	plug :match
	plug :dispatch

	get "/users" do 
		send_resp(conn, 200, "get users")
	end

	get "/users/:id" do
		send_resp(conn, 200, "get user #{id}")
	end

	post "/users" do 
		send_resp(conn, 201, "create user")
	end

	put "/users/:id" do 
		send_resp(conn, 200, "update user #{id}")
	end

	patch "/users/:id" do 
		send_resp(conn, 200, "delete user #{id}")
	end

	get "/reviews" do 
		send_resp(conn, 200, "get reviews")
	end

	get "/reviews/:id" do 
		send_resp(conn, 200, "get review #{id}")
	end

	post "/reviews" do 
		send_resp(conn, 201, "create review")
	end

	put "/reviews/:id" do 
		send_resp(conn, 200, "update review #{id}")
	end

	patch "/reviews/:id" do 
		send_resp(conn, 200, "delete review #{id}")
	end

	get "/venues" do 
		send_resp(conn, 200, "get venues")
	end

	get "/venues/:id" do 
		send_resp(conn, 200, "get venue #{id}")
	end

	post "/venues" do 
		send_resp(conn, 201, "create venue")
	end

	put "/venues/:id" do 
		send_resp(conn, 200, "update venue #{id}")
	end

	patch "/venues/:id" do 
		send_resp(conn, 200, "delete venue #{id}")
	end

	get "/achievements" do 
		send_resp(conn, 200, "get achievements")
	end

	get "/achievements/:id" do 
		send_resp(conn, 200, "get achievement #{id}")
	end

	post "/achievements" do 
		send_resp(conn, 201, "create achievement")
	end

	put "/achievements/:id" do 
		send_resp(conn, 200, "update achievement #{id}")
	end

	delete "/achievements/:id" do 
		send_resp(conn, 204, "delete achievement #{id}")
	end

	# match "/users/*_rest" do
	#	conn |> IO.inspect
	#	{conn, status, body} = conn |> handle_request
	#	send_resp(conn, status, body)
	# end

	#def handle_request(%Plug.Conn{
	#	method: method,
	#	path_info: path_info
	#} = conn) when is_get_all(method, path_info) do
	#	{conn, 200, "wew lad"}
	#end

	# forward "/users", to: Parmyay.Routers.Users

	match _, do: send_resp(
		conn,
		404,
		"These are not the droids you are looking for"
	)

	defp handle_errors(conn, %{kind: _kind, reason: _reason, stack: _stack}) do
    	send_resp(conn, conn.status, "Something went wrong")
  	end
end
