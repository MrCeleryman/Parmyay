defmodule Parmyay.Routers.Reviews do
  	use Plug.Router

	if Mix.env == :dev do
    	use Plug.Debugger
  	end

	use Plug.ErrorHandler
	
	plug :match
	plug :dispatch

	get "/" do
    	send_resp(conn, 200, "get all")
  	end

	get "/:id" do
    	send_resp(conn, 200, "get #{id}")
  	end

	post "/" do
    	send_resp(conn, 200, "post")
  	end

	put "/:id" do
    	send_resp(conn, 200, "update #{id}")
  	end

	patch "/:id" do
    	send_resp(conn, 200, "delete #{id}")
  	end

	match _, do: send_resp(conn, 404, "Oops!")
end
