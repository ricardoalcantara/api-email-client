import { LocalStorageHook, useLocalStorage } from "@/hooks/useLocalStorage";
import { useEffect } from "react";

export enum Theme {
  Dark = "dark",
  Light = "light",
}

/**
 * Theme hook
 */
export const useTheme = (): LocalStorageHook<Theme> => {
  const [theme, setTheme] = useLocalStorage<Theme>("theme", Theme.Light);

  /**
   * Adds or removes dark mode of body tag when theme in local storage changes
   */
  useEffect(() => {
    const bodyClass = window.document.body.classList;

    if (theme === Theme.Dark) {
      bodyClass.add(Theme.Dark);
    } else {
      bodyClass.remove(Theme.Dark);
    }
  }, [theme]);

  return [theme, setTheme];
};
