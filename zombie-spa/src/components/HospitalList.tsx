import React, { FunctionComponent } from 'react';

import { Hospital } from '../models/hospitals';
import { Link } from 'react-router-dom';

interface HospitalListProps {
  items: Hospital[];
}

interface HospitalListItemProps {
  hospital: Hospital;
}

const level = 0;

const HostpitalListItem: FunctionComponent<HospitalListItemProps> = ({
  hospital,
}) => {
  const waitTimeMinutes =
    hospital.waitingList[level].averageProcessTime *
    hospital.waitingList[level].patientCount;

  return (
    <Link to="/">
      <li>
        <span>{hospital.name}</span>
        <span>Wait time: {waitTimeMinutes} mins</span>
      </li>
    </Link>
  );
};

const HospitalList: FunctionComponent<HospitalListProps> = (props) => {
  return (
    <ul>
      <p>Our suggested Hospitals:</p>
      {props.items
        .sort((h1, h2) => {
          const averageH1 =
            h1.waitingList[level].averageProcessTime *
            h1.waitingList[level].patientCount;

          const averageH2 =
            h2.waitingList[level].averageProcessTime *
            h2.waitingList[level].patientCount;

          if (averageH1 < averageH2) {
            return -1;
          } else {
            return 0;
          }
        })
        .map((item) => (
          <HostpitalListItem key={`illness-${item.id}`} hospital={item} />
        ))}
    </ul>
  );
};

export default HospitalList;
