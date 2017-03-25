import {Vue, Component, Watch} from "av-ts";
import {LocalisedStrings} from "../../../util/localisedStrings";

@Component({
	template: (
		`<main class="mdl-layout__content">
			<div class="page-content">
				<div class="main-text">
					<h1>Parmyay</h1>
					<h2>Here to help you find your nearest {{ localisations.parma }}</h2>
				</div>
			</div>
		</main>`
	)
})
export default class FrontPage extends Vue {
	public localisations = LocalisedStrings.currentLocalisation;
}
