import { Vue, Component } from "av-ts";
import Switch from "../switch";

@Component({
	components: {
		"parmy-switch": Switch
	},
	template: (
		`<header class="mdl-layout__header">
			<div class="mdl-layout__header-row">
				<router-link :to="{ name: 'home'}">
					<span class="mdl-layout-title">Parmyay</span>
				</router-link>
				<router-link :to="{ name: 'test'}" class="testPageLink">Test Page</router-link>
				<div class="mdl-layout-spacer"></div>
				<parmy-switch></parmy-switch>
			</div>
		</header>`
	)
})
export default class NavBar extends Vue {
}
