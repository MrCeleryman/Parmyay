import { Vue, Component } from "av-ts";
import Switch from "../switch";

@Component({
	components: {
		"parmy-switch": Switch
	},
	template: require("./navbar.html")
})
export default class NavBar extends Vue {
}
