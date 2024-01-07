import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';

import { Deal } from '../models/deal';
@Component({
  selector: 'app-deal',
  templateUrl: './deal.component.html',
  styleUrl: './deal.component.css'
})
export class DealComponent implements OnInit {

  @Input() deal: Deal | undefined
  ngOnInit(): void {
  }


}
