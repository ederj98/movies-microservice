import styled from 'styled-components';

export const Table = styled.table`
  &&& {
    table {
      width: 100%;
      border: 1px solid #999;
      text-align: left;
      border-collapse: collapse;
      margin: 0 0 1em 0;
      caption-side: top;
   }
   td, th {
      padding: 0.3em;
   }
   th, td {
      border-bottom: 1px solid #999;
      width: 25%;
   }
  }
`;
