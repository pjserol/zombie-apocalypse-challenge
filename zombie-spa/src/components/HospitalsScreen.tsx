import React, { FunctionComponent, useEffect, useState } from 'react';
import { HospitalResponse, Hospital } from '../models/hospitals';
import HospitalList from './HospitalList';
import { API_URL } from '../const';

const HospitalsScreen: FunctionComponent = () => {
  const [page, setPage] = useState(0);
  const [hospitals, setHospitals] = useState<Hospital[]>([]);
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    fetch(`${API_URL}hospitals?limit=10&page=${page}`, {
      method: 'GET',
      headers: new Headers({
        Accept: 'application/hal+json;charset=UTF-8',
      }),
    })
      .then((res) => res.json())
      .then(({ _embedded }: HospitalResponse) => {
        const { hospitals } = _embedded;
        setHospitals(hospitals);
        setIsLoaded(true);
      })
      .catch((error) => {
        setIsLoaded(true);
        console.error(error);
      });
  }, [page]);

  return (
    <React.Fragment>
      <HospitalList items={hospitals} />
    </React.Fragment>
  );
};

export default HospitalsScreen;
