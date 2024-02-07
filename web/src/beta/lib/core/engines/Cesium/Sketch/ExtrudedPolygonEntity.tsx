import {
  CallbackProperty,
  ClassificationType,
  ColorMaterialProperty,
  HeightReference,
  ShadowMode,
  type Color,
  type PolygonHierarchy,
} from "@cesium/engine";
import { useMemo, useRef, type FC } from "react";

import { useConstant } from "@reearth/beta/utils/util";

import { useContext } from "../Feature/context";

import { Entity, type EntityProps } from "./Entity";

export interface ExtrudedPolygonEntityProps {
  dynamic?: boolean;
  id?: string;
  hierarchy: PolygonHierarchy;
  extrudedHeight: number;
  color?: Color;
  disableShadow?: boolean;
}

export const ExtrudedPolygonEntity: FC<ExtrudedPolygonEntityProps> = ({
  dynamic = false,
  id,
  hierarchy: hierarchyProp,
  extrudedHeight: extrudedHeightProp,
  color,
  disableShadow = false,
}) => {
  const hierarchyRef = useRef(hierarchyProp);
  hierarchyRef.current = hierarchyProp;
  const hierarchyProperty = useConstant(
    () => new CallbackProperty(() => hierarchyRef.current, false),
  );
  const hierarchy = dynamic ? hierarchyProperty : hierarchyProp;

  const extrudedHeightRef = useRef(extrudedHeightProp);
  extrudedHeightRef.current = extrudedHeightProp;
  const extrudedHeightProperty = useConstant(
    () => new CallbackProperty(() => extrudedHeightRef.current, false),
  );
  const extrudedHeight = dynamic ? extrudedHeightProperty : extrudedHeightProp;

  const options = useMemo(
    (): EntityProps => ({
      polygon: {
        hierarchy,
        extrudedHeight,
        extrudedHeightReference: HeightReference.RELATIVE_TO_GROUND,
        fill: true,
        material: new ColorMaterialProperty(color),
        classificationType: ClassificationType.TERRAIN,
        shadows: disableShadow ? ShadowMode.DISABLED : ShadowMode.ENABLED,
      },
    }),
    [extrudedHeight, disableShadow, hierarchy, color],
  );

  const { requestRender } = useContext();
  requestRender?.();

  return <Entity id={id} {...options} />;
};