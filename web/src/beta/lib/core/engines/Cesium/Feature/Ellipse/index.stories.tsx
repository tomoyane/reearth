import { Meta, Story } from "@storybook/react";

import { engine } from "../..";
import Component, { Props } from "../../../../Map";

export default {
  component: Component,
  parameters: { actions: { argTypesRegex: "^on.*" } },
} as Meta;

const Template: Story<Props> = args => <Component {...args} />;

export const Default = Template.bind([]);
Default.args = {
  engine: "cesium",
  engines: {
    cesium: engine,
  },
  ready: true,
  layers: [
    {
      id: "l",
      type: "simple",
      data: {
        type: "geojson",
        value: {
          type: "Feature",
          geometry: {
            type: "Point",
            coordinates: [0, 0, 1000],
          },
        },
      },
      ellipse: {
        radius: 1000,
        fillColor: "#FF0000",
      },
    },
  ],
  property: {
    tiles: [
      {
        id: "default",
        tile_type: "default",
      },
    ],
  },
};
