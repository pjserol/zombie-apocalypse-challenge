export interface IllnessResponse {
  _embedded: EmbeddedIllness;
  _links: Links;
  page: Page;
}

export interface EmbeddedIllness {
  illnesses: IllnessData[];
}

export interface IllnessData {
  illness: Illness;
}

export interface Illness {
  name: string;
  id: number;
}

export interface Links {
  next: Next;
}

export interface Next {
  href: string;
}

export interface Page {
  size: number;
  totalElements: number;
  totalPages: number;
  number: number;
}
