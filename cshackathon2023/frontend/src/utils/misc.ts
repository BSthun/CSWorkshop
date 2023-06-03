import { ClassValue, clsx } from "clsx";
import React from "react";
import { StoreApi, UseBoundStore } from "zustand";
import { shallow } from "zustand/shallow";

export interface Props {
  children?: React.ReactNode;
  className?: string;
}

export function cls(...classNames: ClassValue[]) {
  return clsx(classNames);
}

export function compareTrim(a: string, b: string) {
  return a.trim() === b.trim();
}

export function getIndexById<T>(arr: T[], id: string | number, key = "id") {
  return arr.findIndex((a) => (a as any)[key] === id);
}

export function withDefault<T>(objects: T, defaults: Partial<T>) {
  return Object.assign({}, defaults, objects);
}

export function convertToFormData<T extends object>(data: T) {
  const formData = new FormData();
  Object.keys(data).forEach((key) => {
    formData.append(key, (data as any)[key]);
  });
  return formData;
}

export function extract<T, K extends keyof T>(
  useTarget: UseBoundStore<StoreApi<T>>,
  keys: K[]
) {
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const data = useTarget((store) => {
    const extractedField: Partial<T> = {};
    for (const k of keys) extractedField[k] = store[k];
    return extractedField;
  }, shallow);
  return data as Pick<T, K>;
}

export function delay(ms: number) {
  return new Promise((res) => setTimeout(res, ms));
}

export async function getImageData(file: File): Promise<{
  w: number;
  h: number;
  size: number;
  ratio: number;
  url: string;
}> {
  return new Promise((res, rej) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      var img = new Image();
      if (!e) return rej("Unknown error!");
      let url = (img.src = e.target!.result as string);
      // URL.revokeObjectURL(url);
      img.onload = function () {
        const size = Math.max(img.width, img.height);
        const w = (img.width = Math.round((100 * img.width) / size));
        const h = (img.height = Math.round((100 * img.height) / size));
        res({
          w,
          h,
          size,
          ratio: Math.round((w * 100) / h) / 100,
          url,
        });
      };
    };
    reader.readAsDataURL(file);
  });
}

export function shareRef(ref: any, shareRef: any) {
  return (e: any) => {
    ref(e);
    shareRef.current = e;
  };
}

export function removeDeplicate<T>(
  arr: T[] = [],
  compareFn: (a: T, b: T) => boolean = (a, b) => a === b
) {
  let _arr = structuredClone(arr);
  arr.forEach((a) => {
    const count = _arr.filter((b) => compareFn(a, b)).length;
    if (count > 1) {
      _arr = _arr.filter((b) => a != b);
    }
  });
  return _arr;
}
