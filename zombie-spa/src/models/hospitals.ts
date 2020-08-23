export interface HospitalResponse {
  _embedded: EmbeddedHospital;
  _links: Links;
  page: Page;
}

export interface EmbeddedHospital {
  hospitals: Hospital[];
}

export interface Hospital {
  id: number;
  name: string;
  waitingList: WaitingList[];
  location: Location;
}

export interface WaitingList {
  patientCount: number;
  levelOfPain: number;
  averageProcessTime: number;
}

export interface Location {
  lat: number;
  lng: number;
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
