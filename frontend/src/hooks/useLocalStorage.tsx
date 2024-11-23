import { useState, useEffect, Dispatch, SetStateAction } from 'react';

export type LocalStorageHook<T> = [T, Dispatch<SetStateAction<T>>];

/**
 * Local storage hook
 * @param key Key
 * @param initialValue Initial Value
 */
export const useLocalStorage = <T extends string>(
  key: string,
  initialValue: string,
): LocalStorageHook<T> => {
  const [storedValue, setStoredValue] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(key);

      return item ? JSON.parse(item) : initialValue;
    } catch (e) {
      return initialValue;
    }
  });

  useEffect(() => {
    try {
      window.localStorage.setItem(key, JSON.stringify(storedValue));
    } catch (e) {
      console.error(e);
    }
  }, [key, storedValue]);

  return [storedValue, setStoredValue];
};
