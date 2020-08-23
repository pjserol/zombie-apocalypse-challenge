import React, { useEffect, useState, FunctionComponent } from 'react';
import IllnessList from './IllnessList';
import { IllnessData, IllnessResponse } from '../models/illnesses';

const IllnessesScreen: FunctionComponent = (props: {}) => {
  const [page, setPage] = useState(0);
  const [illnesses, setIllnesses] = useState<IllnessData[]>([]);
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    fetch(
      `http://dmmw-api.australiaeast.cloudapp.azure.com:8080/illnesses?limit=10&page=${page}`,
      {
        method: 'GET',
        headers: new Headers({
          Accept: 'application/hal+json;charset=UTF-8',
        }),
      },
    )
      .then((res) => res.json())
      .then(({ _embedded }: IllnessResponse) => {
        const { illnesses } = _embedded;
        setIllnesses(illnesses);
        setIsLoaded(true);
      })
      .catch((error) => {
        setIsLoaded(true);
        console.log(error);
      });
  }, [page]);

  return (
    <React.Fragment>
      <IllnessList items={illnesses} />
    </React.Fragment>
  );
};

export default IllnessesScreen;
