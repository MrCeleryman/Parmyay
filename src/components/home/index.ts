import { Vue, Component } from "av-ts";
import NavBar from "../navBar";
import FrontPage from "./frontPage";

@Component({
	components: {
		"front-page": FrontPage,
		"nav-bar": NavBar,
	},
	template: (
		`<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
			<nav-bar></nav-bar>
			<router-view></router-view>
		</div>`
	)
})
export default class Main extends Vue {
}
