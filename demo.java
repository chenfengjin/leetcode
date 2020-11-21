



vector<double> randPoint() {
    double x0 = xc - rad;
    double y0 = yc - rad;

    while(true) {
    double xg = x0 + uni(rng) * 2 * rad;
    double yg = y0 + uni(rng) * 2 * rad;
    if (sqrt(pow((xg - xc), 2) + 
    pow((yg - yc), 2)) <= rad)
        return {xg, yg};
    }
}


