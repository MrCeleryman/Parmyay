import {Vue, Component, Prop} from "av-ts";

@Component({
	template: (
		`<main class="mdl-layout__content">
			<div class="page-content">
				<div class="main-text">
					<h1>Test Page</h1>
					<h2>This page is to test the routing</h2>
				</div>
			</div>
		</main>`
	)
})
export default class TestPage extends Vue {
	@Prop
	public LANG;
}
