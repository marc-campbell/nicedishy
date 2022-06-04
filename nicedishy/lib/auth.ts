import { loadSession } from "./session";

export async function isAuthed(token: string): Promise<boolean> {
  const sess = await loadSession(token);
  if (!sess) {
    return false;
  }

  return true;
}
