import { Vue, Component } from "av-ts";
import NavBar from "../navBar";
import FrontPage from "./frontPage";

@Component({
	components: {
		"front-page": FrontPage,
		"nav-bar": NavBar,
	},
	template: require("./home.html")
})
export default class Main extends Vue {
}
