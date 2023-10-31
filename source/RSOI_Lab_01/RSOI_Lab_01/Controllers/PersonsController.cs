using Microsoft.AspNetCore.Mvc;
using RSOI_Lab_01.Interfaces;
using RSOI_Lab_01.Models;

namespace RSOI_Lab_01.Controllers
{
    [Route("api/v1/persons")]
    [ApiController]
    public class PersonsController : Controller
    {
        readonly IPersonsService _personsService;

        public PersonsController(IPersonsService personsService)
        {
            _personsService = personsService;
        }

        /// <summary>
        /// Get Person by ID
        /// </summary>
        /// <param name="id"></param>
        /// <returns>Person for ID</returns>
        /// <response code="200">Person for ID</response>
        /// <response code="404">Not found Person for ID</response>
        [HttpGet("{id}")]
        [Produces("application/json")]
        [ProducesResponseType(200, Type = typeof(PersonResponse))]
        [ProducesResponseType(404, Type = typeof(ErrorResponse))]
        public async Task<IActionResult> GetPerson(int id)
        {
            var res = await _personsService.GetAsync(id);

            return res is null ? NotFound(new ErrorResponse($"Person with id = {id} was not found")) : Ok(res);
        }

        /// <summary>
        /// Get all Persons
        /// </summary>
        /// <returns>All Persons</returns>
        /// <response code="200">All Persons</response>
        [HttpGet]
        [Produces("application/json")]
        [ProducesResponseType(200, Type = typeof(List<PersonResponse>))]
        public async Task<IActionResult> GetPersons()
        {
            return Ok(await _personsService.GetAllAsync());
        }

        /// <summary>
        /// Create new Person
        /// </summary>
        /// <param name="person"></param>
        /// <returns>Created new Person</returns>
        /// <response code="201">Created new Person</response>
        /// <response code="400">Invalid data</response>
        [HttpPost]
        [Produces("application/json")]
        [ProducesResponseType(400, Type = typeof(ValidationErrorResponse))]
        public async Task<IActionResult> CreatePerson([FromBody] PersonRequest person)
        {
            try
            {
                var ans = await _personsService.CreateAsync(person);

                return Created($"api/v1/persons/{ans!.Value}", null);
            }
            catch (Exception e)
            {
                return BadRequest(e.Message);
            }
        }

        /// <summary>
        /// Update Person by ID
        /// </summary>
        /// <param name="id"></param>
        /// <param name="person"></param>
        /// <returns></returns>
        /// <response code="200">Person for ID was updated</response>
        /// <response code="400">Invalid data</response>
        /// <response code="404">Not found Person for ID</response>
        [HttpPatch("{id}")]
        [Produces("application/json")]
        [ProducesResponseType(200, Type = typeof(PersonResponse))]
        [ProducesResponseType(400, Type = typeof(ValidationErrorResponse))]
        [ProducesResponseType(404, Type = typeof(ErrorResponse))]
        public async Task<IActionResult> UpdatePerson(int id, [FromBody] PersonRequest person)
        {
            try
            {
                var ans = await _personsService.UpdateAsync(id, person);

                return ans is not null ? Ok(ans) : NotFound(new ErrorResponse($"Person with id = {id} was not found"));
            }
            catch (Exception e)
            {
                return BadRequest(e.Message);
            }
        }

        /// <summary>
        /// Remove Person by ID
        /// </summary>
        /// <param name="id"></param>
        /// <returns>Person for ID was removed</returns>
        /// <response code="204">Person for ID was removed</response>
        [HttpDelete("{id}")]
        [ProducesResponseType(204)]
        public async Task<IActionResult> DeletePerson(int id)
        {
            await _personsService.DeleteAsync(id);

            return NoContent();
        }
    }
}