import { Vue, Component, Lifecycle } from "av-ts";
import { AllowedLocalisations, LocalisedStrings } from "../../util/localisedStrings";
import * as d3 from "d3";

@Component({
	template: (
		`<svg :width="wheelDimensions.width" :height="wheelDimensions.height/2" @click="changeOption">
			<g :transform="'translate(' + wheelDimensions.width/2 + ', ' + wheelDimensions.height/4 + ')'">
				<g 
					class="indicator"
					v-for="noun in localisations"
					v-bind:style="{transform: 'rotate(' + noun.rotation + 'deg)', visibility: noun.visibility}">
					<text>{{ noun.text }}</text>
					<line :x1="noun.x1" :x2="noun.x2" :y1="noun.y1" :y2="noun.y2"></line>
				</g>
			</g>
		</svg>`
	)
})
export default class Switch extends Vue {
	private changeIndex = 0;
	private localisations = [];

	private wheelDimensions = {
		height: 0,
		width: 0
	};

	private afterInitialise(noun) {
		noun.y1 = 0;
		noun.y2 = this.wheelDimensions.height / 2 - this.wheelDimensions.height / 4;
		noun.visibility = "visible";
	}

	private calculateRotation() {
		const rotateOut = this.localisations[this.changeIndex];
		this.afterInitialise(rotateOut);
		rotateOut.rotation -= 180;

		this.changeIndex = ((++this.changeIndex) % this.localisations.length);
		const newLocalisation = (this.localisations[this.changeIndex].text as AllowedLocalisations);

		const rotateIn = this.localisations[this.changeIndex];
		this.afterInitialise(rotateIn);
		rotateIn.rotation -= 180;

		LocalisedStrings.setLocalisation(newLocalisation);
		this.$emit("changeOption", newLocalisation);
	}

	@Lifecycle
	private mounted() {
		const dimensions = (this.$el.parentNode as HTMLElement).getBoundingClientRect();
		this.wheelDimensions.height = dimensions.height;
		this.wheelDimensions.width = dimensions.height;

		this.localisations = LocalisedStrings.getLocalisations().map((x) => {
			return {
				"rotation": 0,
				"text": x,
				"visibility": "hidden",
				"x1": 0,
				"x2": 0,
				"y1": 0,
				"y2": (dimensions.height / 2 - (dimensions.height / 4))
			};
		});

		// Strange issue where the css rotation-origin is not applied to the elements.
		// Applying a setTimeout resolves it since the item is rendered, then changed.
		// Does not affect the UX majorly since it is invisible when all this happens.
		setTimeout(() => {
			this.localisations.forEach((x, i) => {
				if (i !== 0) { x.rotation = 180; }
			});
		}, 100);
		this.localisations[0].rotation = 0;
		this.localisations[0].visibility = "visible";
	}

	private changeOption(): void {
		this.calculateRotation();
		/*
		*/
	}
}
